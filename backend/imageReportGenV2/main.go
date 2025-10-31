package main

import (
	"context"
	"embed"
	"log"

	"imageReportGenV2/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	irap := app.NewImageReportApp()

	if err := wails.Run(&options.App{
		Title:  "ImageReportGen V2",
		Width:  960,
		Height: 640,
		OnStartup: func(ctx context.Context) {
			irap.Startup(ctx)
		},
		AssetServer:      &assetserver.Options{Assets: assets},
		Bind:             []interface{}{irap},
		BackgroundColour: &options.RGBA{R: 13, G: 17, B: 23, A: 1},
	}); err != nil {
		log.Fatal(err)
	}
}
