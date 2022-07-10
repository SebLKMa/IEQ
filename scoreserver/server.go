package main

import (
	"log"
	"sync"

	tasks "github.com/seblkma/ieq/tasks"
)

func doAwairScores() {
	defer wg.Done()
	task := tasks.NewScoringTask("configawair.yaml")
	task.Execute()
}

func doUhooScores() {
	defer wg.Done()
	task := tasks.NewScoringTask("configuhoo.yaml")
	task.Execute()
}

// for keeping main() running
var wg = &sync.WaitGroup{}

// You must build and run as binary, pkill or control-c to quit
// go build -o server server.go
func main() {

	// Never run Empty loop in main, empty loop uses 100% of a CPU Core
	// https://stackoverflow.com/questions/39902477/whats-the-best-practices-to-run-a-background-task-along-with-server-listening
	// https://stackoverflow.com/questions/39493692/difference-between-the-main-goroutine-and-spawned-goroutines-of-a-go-program

	wg.Add(1)
	go doAwairScores()
	log.Println("awair task started")
	wg.Add(1)
	go doUhooScores()
	log.Println("uhoo task started")

	wg.Wait() // blocks until all tasks are completed or our pid is killed
}
