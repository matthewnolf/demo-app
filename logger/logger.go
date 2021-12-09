package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/go-kit/log"
)

const depth = 3

var callerFunc = func(repoName string, revision string) log.Valuer {
	return func() interface{} {
		_, file, line, _ := runtime.Caller(depth)

		parts := strings.Split(file, strings.TrimPrefix(repoName, "github.com/matthewnolf/"))
		path := parts[len(parts)-1]

		return fmt.Sprintf("%s;%s;%s;%d", repoName, revision, path, line)
	}
}

func Logger(repoURL string, revision string) log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "caller", callerFunc(repoURL, revision))

	return logger
}
