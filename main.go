package main

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/arrim/go-cart/app/repository"
	"github.com/arrim/go-cart/app/repository/cart"
	"github.com/arrim/go-cart/server"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

type Config struct {
	Server server.Options
	DBUrl  string
}

var configFile = flag.String("config", "config/test.toml", "Path to configuration file")

func main() {
	flag.Parse()

	config := loadConfig()

	conn, err := pgx.Connect(context.Background(), config.DBUrl)

	if err != nil {
		log.Printf("Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	repo := repository.NewRepository(cart.NewPostgresCartRepo(conn))
	g := server.NewServer(config.Server, repo)
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
