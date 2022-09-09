package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once     = &sync.Once{}
	instance *logrus.Logger
)

func Get() *logrus.Logger {
	once.Do(func() {
		logger := logrus.New()

		logger.SetReportCaller(true)
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrus.TraceLevel)

		logger.Formatter = &logrus.TextFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				filename := path.Base(frame.File)
				return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
			},
			DisableColors: false,
			FullTimestamp: false,
		}

		instance = logger
	})

	return instance
}
