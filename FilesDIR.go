//go:generate goversioninfo -icon=FilesDIR.ico -manifest=FilesDIR.exe.manifest
package main

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"FilesDIR/pkg"
	"FilesDIR/rgb"
	"bufio"
	"flag"
	"os"
	"time"
)

func main() {

	// Flag of Packages
	FlgCls := flag.Bool("cls", false, "Nettoie les dossiers logs, dumps et exports")
	FlgCompiler := flag.Bool("c", false, "Lance le mode de compilation")
	// Flag of search
	FlgMode := flag.String("mode", "%", "Mode de recherche")
	FlgWord := flag.String("word", "", "Non de fichier")
	FlgExt := flag.String("ext", "*", "Ext de fichier")
	FlgPoolSize := flag.Int("poolsize", 10, "Nombre de tâches en simultanées")
	// Flag of criteral of search
	FlgMaj := flag.Bool("maj", false, "Autorise les majuscules")
	// Flag of special mode
	FlgDevil := flag.Bool("devil", false, "Mode 'Démon' de l'application")
	FlgSilent := flag.Bool("s", false, "Mode 'Silent', évite toutes les choses inutiles")
	FlgBlackList := flag.Bool("b", false, "Ajout d'une blacklist de dossier")
	FlgWhiteList := flag.Bool("w", false, "Ajout d'une whitelist de dossier")
	// Parse all Flags
	flag.Parse()

	s := &pkg.Search{
		Cls:       *FlgCls,
		Compiler:  *FlgCompiler,
		Mode:      *FlgMode,
		Word:      *FlgWord,
		Ext:       *FlgExt,
		PoolSize:  *FlgPoolSize,
		Maj:       *FlgMaj,
		Devil:     *FlgDevil,
		Silent:    *FlgSilent,
		BlackList: *FlgBlackList,
		WhiteList: *FlgWhiteList,

		SrcPath: pkg.GetCurrentDir(),
		DstPath: config.DstPath,
		Timer: &pkg.Timer{
			AppStart: time.Now(),
		},
		Counter: &pkg.Counter{},
		Process: &pkg.Process{},
	}

	s.DrawStart()

	if s.Cls {
		pkg.CleenTempFiles()
	} else if s.Compiler {

	} else {
		s.RunSearch()
	}

	s.Timer.AppEnd = time.Since(s.Timer.AppStart)

	if s.Cls {
		s.DrawCls()
	} else if s.Compiler {

	} else {
		s.DrawBilanSearch()
	}

	s.DrawEnd()

	rgb.GreenB.Print("Appuyer sur Entrée pour quitter...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		loger.Crash("Crash :", err)
	}
}
