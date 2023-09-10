package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"goblinBoard/internal/app"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	title  = "GoblinBoard"
	width  = 1024
	height = 768
)

func main() {
	// Create an instance of the app structure
	NewAppInstance := app.NewApp()

	err := app.Run(&NewAppInstance{
		wailsConfig.Title:  title,
		wailsConfig.Width:  width,
		wailsConfig.Height: height,
		wailsConfig.AssetServer: &assetserver.Options{
			Assets: assets,
		},
		wailsConfig.BackgroundColour & options.RGBA{R: 27, G: 38, B: 54, A: 1},
		wailsConfig.OnStartup: app.startup,
		wailsConfig.Bind: []interface{}{
			app.App,
		},
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "GoblinBoard",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
