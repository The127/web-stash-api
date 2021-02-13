package main

import (
	"flag"
	_ "github.com/lib/pq"
	"web-stash-api/config"
	"web-stash-api/database"
	"web-stash-api/server"
)

func main() {
	configPath := flag.String("config", "config.yml", "the location of the configuration file")
	envPrefix := flag.String("prefix", "", "environment variable prefix")
	flag.Parse()

	config.ReadConfig(*configPath, *envPrefix)
	database.Migrate()
	server.Init()
}
