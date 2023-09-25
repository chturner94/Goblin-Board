package main

import (
	"embed"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/chturner94/Goblin-Board/internal/app"
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
	newAppInstance := &app.App{}

	err := app.Run(&app.App{
		WailsConfig: options.App{
			Title:  title,
			Width:  width,
			Height: height,
			AssetServer: &assetserver.Options{
				Assets: assets,
			},
			BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
			LogLevel:         logger.DEBUG,
			OnStartup:        newAppInstance.Startup,
			Bind: []interface{}{
				newAppInstance,
			},
		},
	})
	// Create application with options
	if err != nil {
		println("Error:", err.Error())
	}
}
