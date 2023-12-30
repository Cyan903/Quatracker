package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/Cyan903/QuaverBuddy/backend/app"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/config"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := app.NewApp()
	config.Data = config.LoadConfig()

	if err := wails.Run(&options.App{
		Title:  "QuaverBuddy",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: app.NewImages(),
		},

		BackgroundColour: &options.RGBA{R: 29, G: 35, B: 42, A: 255},
		OnStartup:        app.Startup,
		Bind:             []interface{}{app},
	}); err != nil {
		log.Error.Println("Fatal: Could not start app")
		log.Error.Println(err.Error())
	}
}
