package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	logrus.SetReportCaller(true)

	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")


	log2 := logrus.New()
	log2.SetLevel(logrus.DebugLevel)
	log2.Formatter = &logrus.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	}

	logger_name := "logrus"
	cur_time := time.Now()
	log_file_name := fmt.Sprintf("%s_%04d-%02d-%02d-%02d-%02d.txt",
		logger_name, cur_time.Year(), cur_time.Month(), cur_time.Day(), cur_time.Hour(), cur_time.Minute())
	log_file, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		fmt.Printf("try create logfile[%s] error[%s]\n", log_file_name, err.Error())
		return
	}

	defer log_file.Close()

	log2.SetOutput(log_file)

	for i := 0; i < 10; i++ {
		log2.Debugf("logrus to file test %d", i)
	}
}
