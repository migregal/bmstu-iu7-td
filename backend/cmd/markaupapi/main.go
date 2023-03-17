package main

import (
	"flag"
	"log"
	"markup2/markaupapi/api"
	"markup2/markaupapi/config"
)

var (
	configPath string
)

func parseFlags() {
	flag.StringVar(&configPath, "config", "/usr/local/etc/markaup2.yaml", "configuration file to use")
	flag.Parse()
}

func main() {
	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	api := api.New(cfg)

	api.Run()
}
