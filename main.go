package main

import (
	"github.com/onurkybsi/rester/app"
	"github.com/onurkybsi/rester/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Init(config)

	app.Run(config.Port)
}
