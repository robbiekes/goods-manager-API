package main

import (
	"github.com/robbiekes/goods-manager-api/config"
	"github.com/robbiekes/goods-manager-api/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.SetLogrus(cfg.Log.Level)

	// Run
	app.Run(cfg)
}
