package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type UTCFormatter struct {
	logrus.Formatter
}

func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.Local()
	return u.Formatter.Format(e)
}

func InitLogrus() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(UTCFormatter{&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			getCaller := func(s int) (f string, m string, l int) {
				p, file, line, _ := runtime.Caller(s)
				dirs := strings.Split(file, "/")
				//dirs = dirs[len(dirs)-3:]
				f = strings.Replace(strings.Trim(fmt.Sprint(dirs), "[]"), " ", "/", -1)
				m = runtime.FuncForPC(p).Name()
				index := strings.Index(m, ".")
				m = m[index+1:]
				m = strings.ReplaceAll(m, ".func1", "")
				l = line
				return
			}
			file, method, line := getCaller(12)
			return method, fmt.Sprintf("%s:%d", file, line)
		},
	}})
	logrus.AddHook(SetFileDriver(cfg.File))
	return
}

func SetFileDriver(fileName string) *lfshook.LfsHook {
	currentPath, _ := os.Getwd()
	fileName = currentPath + "/" + fileName
	logWriter, _ := rotatelogs.New(
		strings.Replace(fileName, ".log", "", -1)+"-%Y-%m-%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			getCaller := func(s int) (f string, m string, l int) {
				p, file, line, _ := runtime.Caller(s)
				dirs := strings.Split(file, "/")
				dirs = dirs[len(dirs)-3:]
				f = strings.Replace(strings.Trim(fmt.Sprint(dirs), "[]"), " ", "/", -1)
				m = runtime.FuncForPC(p).Name()
				index := strings.Index(m, ".")
				m = m[index+1:]
				m = strings.ReplaceAll(m, ".func1", "")
				l = line
				return
			}
			file, method, line := getCaller(13)
			return method, fmt.Sprintf("%s:%d", file, line)
		},
	})
	log.SetOutput(logWriter)
	return lfHook
}
