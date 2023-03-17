package main

import (
	"flag"
	"log"
	"markup2/markupapi/api"
	"markup2/markupapi/config"
)

var (
	configPath string
)

func parseFlags() {
	flag.StringVar(&configPath, "config", "/usr/local/etc/markaup2.yaml", "configuration file to use")
	flag.Parse()
}

func main() {
	parseFlags()

	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	api := api.New(cfg)

	api.Run()
}
