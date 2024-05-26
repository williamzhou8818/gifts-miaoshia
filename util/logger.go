package util

import (
	"fmt"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	// 在《I/O进阶编程》(https://www.bilibili.com/cheese/play/ss5925)这门课程里自己实现了一个logger，功能齐全，性能比LogRus快10倍。
	LogRus *logrus.Logger
)

func InitLog(configFile string) {
	viper := CreateConfig(configFile)
	LogRus = logrus.New()
	switch strings.ToLower(viper.GetString("level")) {
	case "debug":
		LogRus.SetLevel(logrus.DebugLevel)
	case "info":
		LogRus.SetLevel(logrus.InfoLevel)
	case "warn":
		LogRus.SetLevel(logrus.WarnLevel)
	case "error":
		LogRus.SetLevel(logrus.ErrorLevel)
	case "panic":
		LogRus.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level %s", viper.GetString("level")))
	}

	LogRus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	logFile := ProjectRootPath + viper.GetString("file")
	fout, err := rotatelogs.New(
		logFile+".%Y%m%d%H",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationTime(1*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	LogRus.SetOutput(fout)
	LogRus.SetReportCaller(true)
}
