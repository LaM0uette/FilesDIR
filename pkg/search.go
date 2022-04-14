package pkg

import (
	"FilesDIR/loger"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Timer struct {
	AppStart time.Time
	AppEnd   time.Duration

	SearchStart time.Time
	SearchEnd   time.Duration
}

type Counter struct {
	NbrFiles    uint64
	NbrAllFiles uint64
}

type Search struct {

	// Flags
	Cls      bool
	Compiler bool

	//..
	Mode      string
	Word      string
	Ext       string
	PoolSize  int
	Maj       bool
	Devil     bool
	Silent    bool
	BlackList bool
	WhiteList bool

	// Search
	SrcPath string
	DstPath string
	ReqUse  string

	// Data
	ListBlackList []string
	ListWhiteList []string
	Timer         *Timer
	Counter       *Counter
}

var (
	wgSch   sync.WaitGroup
	jobsSch = make(chan string)
)

//...
// Functions
func (s *Search) RunSearch() {
	s.initSearch()

	s.Timer.SearchStart = time.Now()

	s.loopDirsWorker(s.SrcPath)

	wgSch.Wait()

	time.Sleep(1 * time.Second)

	s.Timer.SearchEnd = time.Since(s.Timer.SearchStart)
}

func (s *Search) initSearch() {
	DrawParam("INITIALISATION DE LA RECHERCHE EN COURS")

	// Construct variable of search
	s.ReqUse = s.getReqOfSearched()
	if !s.Maj {
		s.Word = StrToLower(s.Word)
	}
	s.Ext = fmt.Sprintf(".%s", s.Ext)

	// Add WhiteList / BlackList
	if s.BlackList {
		blPath := filepath.Join(s.DstPath, "blacklist")
		s.setBlackWhiteList(filepath.Join(blPath, "__ALL__.txt"), 0)

		file := filepath.Join(blPath, fmt.Sprintf("%s.txt", StrToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.setBlackWhiteList(file, 0)
		}

		DrawParam(fmt.Sprintf("BLACKLIST: %v", s.ListBlackList))
	}
	if s.WhiteList {
		wlPath := filepath.Join(s.DstPath, "whitelist")
		s.setBlackWhiteList(filepath.Join(wlPath, "__ALL__.txt"), 1)

		file := filepath.Join(wlPath, fmt.Sprintf("%s.txt", StrToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.setBlackWhiteList(file, 1)
		}

		DrawParam(fmt.Sprintf("WHITELIST: %v", s.ListWhiteList))
	}

	// Check basics configurations
	s.checkMinimumPoolSize()
	s.setMaxThread()

	// Creation of workers for search
	for w := 1; w <= s.PoolSize; w++ {
		numWorker := w
		go func() {
			err := s.loopFilesWorker()
			if err != nil {
				loger.Error(fmt.Sprintf("Error with worker N°%v", numWorker), err)
			}
		}()
	}

	// Create csv dump
	loger.Semicolon("id;Fichier;Date;Lien_Fichier;Lien")
}

func (s *Search) getReqOfSearched() string {

	req := "FilesDIR"

	if !s.Cls && !s.Compiler {
		req += fmt.Sprintf(" -mode=%s", s.Mode)

		if s.Word != "" {
			req += fmt.Sprintf(" -word=%s", s.Word)
		}

		if s.Ext != "" {
			req += fmt.Sprintf(" -ext=%s", s.Ext)
		}

		req += fmt.Sprintf(" -poolsize=%v", s.PoolSize)

		if s.Maj {
			req += " -maj"
		}

		if s.Devil {
			req += " -devil"
		}

		if s.BlackList {
			req += " -b"
		}

		if s.WhiteList {
			req += " -w"
		}
	}

	if s.Cls {
		req += " -cls"
	} else if s.Compiler {
		req += " -c"
	}

	if s.Silent {
		req += " -s"
	}

	DrawParam(fmt.Sprintf("REQUETE UTILISEE: %s", req))

	return req
}

func (s *Search) setBlackWhiteList(file string, val int) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err) //TODO: Loger crash
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		switch val {
		case 0:
			s.ListBlackList = append(s.ListBlackList, fileScanner.Text())
		case 1:
			s.ListWhiteList = append(s.ListWhiteList, fileScanner.Text())
		}

	}
	_ = readFile.Close()
}

func (s *Search) checkMinimumPoolSize() {
	if s.PoolSize < 2 {
		s.PoolSize = 2
		DrawParam("POOLSIZE MISE A", strconv.Itoa(s.PoolSize), "(ne peut pas être inférieur)")
	} else {
		DrawParam("POOLSIZE MISE A", strconv.Itoa(s.PoolSize))
	}
}

func (s *Search) setMaxThread() {
	maxThr := s.PoolSize * 500
	debug.SetMaxThreads(maxThr)
	DrawParam("THREADS MIS A", strconv.Itoa(maxThr))
}

func (s *Search) isInBlackList(f string) bool {
	for _, word := range s.ListBlackList {
		if strings.Contains(StrToLower(f), StrToLower(word)) {
			return true
		}
	}
	return false
}

func (s *Search) isInWhiteList(f string) bool {
	for _, word := range s.ListWhiteList {
		if strings.Contains(StrToLower(f), StrToLower(word)) {
			return true
		}
	}
	return false
}

func (s *Search) checkFileSearched(file string) bool {
	name := file[:strings.LastIndex(file, path.Ext(file))]
	ext := StrToLower(filepath.Ext(file))

	if !s.Maj {
		name = StrToLower(name)
	}

	// condition of search Mode ( = | % | ^ | $ )
	switch s.Mode {
	case "%":
		if !strings.Contains(name, s.Word) {
			return false
		}
	case "=":
		if name != s.Word {
			return false
		}
	case "^":
		if !strings.HasPrefix(name, s.Word) {
			return false
		}
	case "$":
		if !strings.HasSuffix(name, s.Word) {
			return false
		}
	default:
		if !strings.Contains(name, s.Word) {
			return false
		}
	}

	// condition of extension file
	if s.Ext != ".*" && ext != s.Ext {
		return false
	}

	// condition of open file
	if strings.Contains(name, "~") {
		return false
	}

	return true
}

//...
// WORKER:
func (s *Search) loopFilesWorker() error {
	for jobPath := range jobsSch {

		files, err := ioutil.ReadDir(jobPath)
		if err != nil {
			loger.Crash("Crash with this path:", err)
			wgSch.Done()
			return err
		}

		for _, file := range files {
			if !file.IsDir() {

				if s.checkFileSearched(file.Name()) {
					fmt.Println(file.Name())
					atomic.AddUint64(&s.Counter.NbrFiles, 1)
					atomic.AddUint64(&s.Counter.NbrAllFiles, 1)
				}

				/*
					if s.checkFileSearched(file.Name()) {
						s.NbFiles++
						s.NbFilesTotal++

						if !super {
							Mu.Lock()
							loger.POOk(display.DrawFileSearched(s.NbFiles, file.Name()))

							dataExp := construct.ExportData{
								Id:       s.NbFiles,
								File:     file.Name(),
								Date:     file.ModTime().Format("02-01-2006 15:04:05"),
								PathFile: filepath.Join(jobPath, file.Name()),
								Path:     jobPath,
							}
							construct.ExcelData = append(construct.ExcelData, dataExp)
							Mu.Unlock()

						} else {
							loger.POAction(display.DrawSearchedFait(s.NbFilesTotal))
						}

						loger.LOOk(fmt.Sprintf("N°%v |=| Files: %s", s.NbFiles, file.Name()))
						loger.Semicolon(fmt.Sprintf("%v;%s;%s;%s;%s",
							s.NbFiles, file.Name(), file.ModTime().Format("02-01-2006 15:04:05"), filepath.Join(jobPath, file.Name()), jobPath))

						if runtime.NumGoroutine() > s.NbGoroutine {
							s.NbGoroutine = runtime.NumGoroutine()
						}
					} else {
						s.NbFilesTotal++
						loger.POAction(display.DrawSearchedFait(s.NbFilesTotal))
					}*/
			}
		}
		wgSch.Done()
	}
	return nil
}

func (s *Search) loopDirsWorker(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		loger.Error("Error with this path:", err)
	}

	go func() {
		wgSch.Add(1)
		jobsSch <- path
	}()

	for _, file := range files {
		if file.IsDir() && !s.isInBlackList(file.Name()) {

			if s.WhiteList && !s.isInWhiteList(file.Name()) {
				return
			}

			if s.Devil {
				time.Sleep(20 * time.Millisecond)
				go s.loopDirsWorker(filepath.Join(path, file.Name()))
			} else {
				s.loopDirsWorker(filepath.Join(path, file.Name()))
			}
		}
	}
}
