package main

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/spf13/pflag"
)

type Config struct {
	ListenAddr string `env:"LISTEN_ADDR"`
}

func NewConfig() *Config {
	c := &Config{}

	pflag.StringVar(&c.ListenAddr, "addr", ":8080", "Listen address")
	pflag.Parse()
	err := env.Parse(c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
