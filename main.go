package main

import (
	"embed"

	"github.com/ethan-mdev/game-launcher/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := backend.NewApp()
	authService := backend.NewAuthService()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "game-launcher",
		Width:     1100,
		Height:    700,
		MinWidth:  900,
		MinHeight: 600,
		MaxWidth:  1100,
		MaxHeight: 700,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
		Bind: []interface{}{
			app,
			authService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
