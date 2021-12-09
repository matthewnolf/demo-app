package main

import (
	"time"

	"github.com/go-kit/log/level"
	"github.com/grafana/hackathon-2021-12-showusthecode/service"
	"github.com/grafana/hackathon-2021-12-showusthecode/worker"

	"github.com/grafana/hackathon-2021-12-showusthecode/logger"
)

func main() {
	for {
		l := logger.Logger("github.com/matthewnolf/demo-app", "main")
		level.Info(l).Log("foo", "bar")

		time.Sleep(2 * time.Second)

		go worker.DoWork(l)
		go service.DoService()
	}
}
