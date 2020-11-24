package logger

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type config struct {
	Enable bool   `env:"LOGGER_ENABLE" envDefault:"false"`
	File   string `env:"LOGGER_FILE" envDefault:"./logs/"`
}

var cfg = &config{}

func init() {
	if err := env.Parse(cfg); err != nil {
		Fatal(err)
	}
	InitLogrus()
	InitZap()
}

func Info(args ...interface{}) {
	if !cfg.Enable {
		return
	}
	logrus.Info(args...)
	//ZapLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if !cfg.Enable {
		return
	}
	//ZapLogger.Infof(template, args...)
	logrus.Infof(template, args...)
}

func Error(args ...interface{}) {
	if !cfg.Enable {
		return
	}
	//ZapLogger.Error(args...)
	logrus.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if !cfg.Enable {
		return
	}
	//ZapLogger.Errorf(template, args...)
	logrus.Errorf(template, args...)
}

func Warn(args ...interface{}) {
	if !cfg.Enable {
		return
	}
	//ZapLogger.Warn(args...)
	logrus.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	if !cfg.Enable {
		return
	}
	//ZapLogger.Warnf(template, args...)
	logrus.Warnf(template, args...)
}

func Panic(args ...interface{}) {
	if !cfg.Enable {
		return
	}
	//ZapLogger.Panic(args...)
	logrus.Panic(args...)
}

func Fatal(args ...interface{}) {
	if !cfg.Enable {
		return
	}
	//ZapLogger.Fatal(args...)
	logrus.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	if !cfg.Enable {
		return
	}
	//ZapLogger.Fatalf(template, args...)
	logrus.Fatalf(template, args...)
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
	return logrus.WithFields(fields)
}
