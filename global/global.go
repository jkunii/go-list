package global

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jkunii/go-list/helper"
	"gopkg.in/caarlos0/env.v2"
)

type Config struct {
	ServerDomain string `env:"SERVER_DOMAIN" envDefault:"http://localhost:1323"`
	LogLevel     int    `env:"LOG_LEVEL" envDefault:"5"`
	Port         string `env:"PORT" envDefault:"1323"`
	ShowBanner   bool   `env:"SHOW_BANNER" envDefault:"false"`
	LogLevelType string `env:"LOG_LEVEL_TYPE" envDefault:"debug"`
	LogFormater  string `env:"LOG_FORMATER" envDefault:"txt"`
	UserName     string `env:"BAISC_USER" envDefault:"123"`
	Secret       string `env:"BASIC_SECRET" envDefault:"123"`
}

var (
	Cfg    Config
	Logger = logrus.New()
)

func (r Config) Init() {

	err := env.Parse(&Cfg)
	if err != nil {
		fmt.Println("Error getting environment variables")
		fmt.Println(err)
	}

	//Logger
	if Cfg.LogFormater == "json" {
		Logger.Formatter = new(logrus.JSONFormatter)
	} else {
		Logger.Formatter = new(logrus.TextFormatter) // default
	}

	Logger.Level, err = logrus.ParseLevel(Cfg.LogLevelType)
	helper.PanicErr(err)

}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Warning(args ...interface{}) {
	Logger.Warning(args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func LogLevel() logrus.Level {
	return Logger.Level
}
