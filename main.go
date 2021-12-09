package main

import (
	"time"

	"github.com/go-kit/log/level"

	"github.com/matthewnolf/demo-app/logger"
	"github.com/matthewnolf/demo-app/service"
	"github.com/matthewnolf/demo-app/worker"
)

var gitRevision = "dev"

func main() {
	for {
		l := logger.Logger("github.com/matthewnolf/demo-app", "main")
		level.Info(l).Log("foo", "bar")

		time.Sleep(2 * time.Second)

		go worker.DoWork(l)
		go worker.ConvertToInteger(l)
		go service.DoService(gitRevision)
	}
}
