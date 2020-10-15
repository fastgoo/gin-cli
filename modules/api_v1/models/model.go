package models

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

type config struct {
	Master    string   `env:"MYSQL_MASTER,required"`
	Slaves    []string `env:"MYSQL_SLAVES" envSeparator:","`
	IdleTime  int      `env:"MYSQL_IDLETIME" envDefault:"3600"`
	Lifetime  int      `env:"MYSQL_LIFTTIME" envDefault:"86400"`
	IdleConns int      `env:"MYSQL_IDLECONNS" envDefault:"20"`
	OpenConns int      `env:"MYSQL_OPENCONNS" envDefault:"100"`
}

func SetUp() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	var err error
	DB, err = gorm.Open(mysql.Open(cfg.Master), &gorm.Config{})
	if err != nil {
		log.Fatal("master: ", err)
	}

	var slaves []gorm.Dialector
	for _, slave := range cfg.Slaves {
		slaves = append(slaves, mysql.Open(slave))
	}
	err = DB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: slaves,
		Policy:   dbresolver.RandomPolicy{},
	}).SetConnMaxIdleTime(time.Duration(cfg.IdleTime) * time.Second).
		SetConnMaxLifetime(time.Duration(cfg.IdleTime) * time.Second).
		SetMaxIdleConns(cfg.IdleConns).
		SetMaxOpenConns(cfg.OpenConns))
	if err != nil {
		log.Fatal("slave: ", err)
	}
}
