package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.InfoLevel)
	Log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/pmx.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	})
}
