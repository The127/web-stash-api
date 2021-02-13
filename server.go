package main

import (
	"context"
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"web-stash-api/config"
	"web-stash-api/ent"
	"web-stash-api/ent/migrate"
)

func main() {
	configPath := flag.String("config", "config.yml", "the location of the configuration file")
	envPrefix := flag.String("prefix", "", "environment variable prefix")

	flag.Parse()

	var cfg config.Config
	readFileConfig(&cfg, *configPath)
	readEnv(&cfg, *envPrefix)

	client, err := ent.Open(cfg.Database.Driver, cfg.Database.Address)
	if err != nil {
		log.Fatalf("failed connecting to database: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	// run migration
	err = client.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
	if err != nil {
		log.Fatalf("failed to create schema resources: %v", err)
	}

	r := gin.Default()

	if len(cfg.Server.Cors.AllowOrigins) > 0 {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = cfg.Server.Cors.AllowOrigins
		r.Use(cors.New(corsConfig))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(getBindAddress(cfg))
}

func getBindAddress(cfg config.Config) string {
	if cfg.Server.BindAddress == "" {
		return "localhost:10000"
	}
	return cfg.Server.BindAddress
}

func readFileConfig(cfg *config.Config, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	} else if err != nil {
		log.Fatalf("error while checking config yaml: %v", err)
	}

	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("error while reading config yaml: %v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}
}

func readEnv(cfg *config.Config, prefix string) {
	err := envconfig.Process(prefix, cfg)
	if err != nil {
		log.Fatalf("failed to process environment variables for configuration: %v", err)
	}
}
