package service

import (
	"time"

	"github.com/go-kit/log/level"
	"github.com/matthewnolf/demo-app/logger"
)

func DoService() {
	for {
		l := logger.Logger("github.com/matthewnolf/demo-app", "b07bc36553512dc443072d805536a119b64dca72")
		level.Info(l).Log("different", "branch")

		time.Sleep(5 * time.Second)
	}
}
