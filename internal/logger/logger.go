package logger

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	logFile := &lumberjack.Logger{
		Filename:   "./logs/pmx.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	Log.SetOutput(multiWriter)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.InfoLevel)
}
