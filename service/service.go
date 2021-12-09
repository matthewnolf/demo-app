package service

import (
	"time"

	"github.com/go-kit/log/level"
	"github.com/matthewnolf/demo-app/logger"
)

func DoService() {
	for {
		l := logger.Logger("github.com/matthewnolf/demo-app", "8f348bb348e35b670169b7fe1a26ac9ff97cb7ba")
		level.Info(l).Log("different", "branch")

		time.Sleep(5 * time.Second)
	}
}
