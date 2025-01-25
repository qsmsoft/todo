package main

import (
	"github.com/qsmsoft/todo/config"
	"github.com/qsmsoft/todo/internal/app"
)

func main() {
	cfg := config.LoadConfig()

	e := app.SetupApp()

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
