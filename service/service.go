package service

import (
	"github.com/go-kit/log/level"
	"github.com/grafana/hackathon-2021-12-showusthecode/logger"
	"time"
)

func DoService() {
	for {
		l := logger.Logger("github.com/grafana/hackathon-2021-12-showusthecode", "348c6fe5220de25fc19157cd07a2c02efcfc5e19")
		level.Info(l).Log("different", "branch")

		time.Sleep(5 * time.Second)
	}
}
