package service

import (
	"time"

	"github.com/go-kit/log/level"
	"github.com/matthewnolf/demo-app/logger"
)

func DoService(gitRevision string) {
	for {
		l := logger.Logger("github.com/matthewnolf/demo-app", gitRevision)
		level.Info(l).Log("log", "target a specific revision")

		time.Sleep(5 * time.Second)
	}
}
