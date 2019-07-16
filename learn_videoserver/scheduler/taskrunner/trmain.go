package taskrunner

import (
	"time"
	// "log"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker {
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <- w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	// start video file cleaning
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()

	// Something else
	// r1 := 
	// w1
	// go w1.
}