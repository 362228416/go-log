package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/op/go-logging"
	"io/ioutil"
	"os"
)

var log = logging.MustGetLogger("go-log")

func init() {
	log.ExtraCalldepth = 1
	var format = logging.MustStringFormatter(
		`%{time:2006/01/02 15:04:05} %{shortfile} [%{level}] %{message}`,
	)
	logging.SetFormatter(format)

	logging.SetBackend(logging.NewLogBackend(&lumberjack.Logger{
		Filename: GetLogsDir() + "/app.log",
		MaxSize:  200,  // megabytes
		Compress: true, // disabled by default
	}, "", 0))
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Warning(args ...interface{}) {
	log.Warning(args)
}

func GetLogsDir() string {
	dir := "./logs"
	found := false
	for i := 0; i < 3; i++ {
		fss, _ := ioutil.ReadDir(".")
		for _, info := range fss {
			if info.Name() == "logs" && info.IsDir() {
				found = true
				parentDir, _ := os.Getwd()
				dir = parentDir + "/logs"
				break
			}
		}
		if !found {
			os.Chdir("..")
		}
	}
	log.Debug("Logs dir", dir)
	return dir
}
