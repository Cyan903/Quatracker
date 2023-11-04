package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/Cyan903/QuaverBuddy/backend/app"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := app.NewApp()

	if err := wails.Run(&options.App{
		Title:            "QuaverBuddy",
		Width:            1024,
		Height:           768,
		AssetServer:      &assetserver.Options{Assets: assets},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		OnStartup:        app.Startup,
		Bind:             []interface{}{app},
	}); err != nil {
		log.Error.Println("Fatal: Could not start app")
		log.Error.Println(err.Error())
	}
}
