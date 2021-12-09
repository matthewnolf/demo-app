package worker

import (
	"time"

	"github.com/go-kit/log"
)

func DoWork(l log.Logger) {
	for {
		l.Log("I am", "Doing work")
		time.Sleep(5 * time.Second)
	}
}
