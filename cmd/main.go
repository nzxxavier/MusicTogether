package main

import (
	"flag"

	"nzx.org/music-together/internal/config"
)

func main() {
	flag.Parse()

	configFile := flag.String("config", "./configs/config.yaml", "Path to config file")
	config.InitConfig(configFile)

	config.StartGinEngine()
}
