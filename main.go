package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/arrim/go-cart/server"
	"log"
	"os"
)

type Config struct {
	Server server.Options
}

var configFile = flag.String("config", "config/test.toml", "Path to configuration file")

func main() {
	flag.Parse()

	config := loadConfig()

	g := server.NewServer(config.Server)
	g.Start()

	g.WaitStop()
}

func loadConfig() *Config {
	var config Config

	if *configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := toml.DecodeFile(*configFile, &config); err != nil {
		log.Fatalln("Cannot read configuration file", err)
	}

	return &config
}
