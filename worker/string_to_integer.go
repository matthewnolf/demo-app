package worker

import (
	"github.com/go-kit/log/level"
	"strconv"
	"time"

	"github.com/go-kit/log"
)

func ConvertToInteger(l log.Logger) {
	for {
		if _, err := strconv.Atoi("oops"); err != nil {
			level.Error(l).Log("error", err)
		}
		time.Sleep(5 * time.Second)
	}
}
