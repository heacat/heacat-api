package logger

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			slash := strings.LastIndex(f.File, "/")
			filename := f.File[slash+1:]
			return "", "[" + filename + ":" + strconv.Itoa(f.Line) + "]"
		},
	})

	Log.SetReportCaller(true)
	Log.SetLevel(logrus.TraceLevel)
}
