package redis

import (
	"context"
	"github.com/caarlos0/env/v6"
	"log"
)

type config struct {
	Addr      []string `env:"REDIS_ADDRS,required" envSeparator:","`
	Password  string   `env:"REDIS_PW" envDefault:""`
	DB        int      `env:"REDIS_DB" envDefault:"0"`
	IsCluster bool
}

var (
	cfg          = &config{}
	ctx          = context.Background()
)

func init() {
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	if len(cfg.Addr) > 1 {
		cfg.IsCluster = true
	}
	Initialize()
}

func Initialize() {
	var err error
	if !cfg.IsCluster {
		err = newClient()
	} else {
		err = newCluster()
	}
	if err != nil {
		log.Fatal("Unable to connect to redis " + err.Error())
	}
}
