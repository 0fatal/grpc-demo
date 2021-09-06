package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Grpc struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	} `toml:"grpc"`
}

var AllConfig Config

func init() {
	_, err := toml.DecodeFile("config/config.toml", &AllConfig)
	if err != nil {
		log.Fatalln(err)
	}
}
