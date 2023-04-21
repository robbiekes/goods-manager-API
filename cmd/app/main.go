package main

import (
	"github.com/robbiekes/goods-manager-api/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
