package main

import (
	"FilesDIR/globals"
	"FilesDIR/log"
	"FilesDIR/task"
	"fmt"
	"time"
)

func main() {

	task.DrawStart()

	log.Blank.Println("===== Starting FilesDIR =====\n")
	timerStart := time.Now()

	s := task.Sch{
		SrcPath:  globals.SrcPathGen,
		PoolSize: 10,
		NbFiles:  0,
	}

	log.Blank.Printf(fmt.Sprintf("===== Starting search on: %s =====\n\n", s.SrcPath))
	task.RunSearch(&s)

	log.Blank.Println("===== Ending search =====\n")
	timerEnd := time.Since(timerStart)

	log.Blank.Println("===== Closing FilesDIR =====")
	task.DrawEnd(&s, timerEnd)
}
