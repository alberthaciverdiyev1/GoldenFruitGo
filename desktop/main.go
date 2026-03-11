package main

import (
	"desktop/internal/handlers"
	"embed"
	"fmt"
	"log"

	apiCmd "github.com/alberthaciverdiyev1/goldenfruit/cmd/api" // Ad toqquşmaması üçün alias
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var appIcon []byte

func main() {
	go apiCmd.Start()

	log.Println("API portu gözlənilir...")
	serverPort := <-apiCmd.PortChan
	apiUrl := fmt.Sprintf("http://127.0.0.1:%d/api/v1", serverPort)
	log.Printf("Frontend API URL: %s", apiUrl)

	// API URL-i App-ə ötürürük
	app := handler.NewApp(apiUrl)

	err := wails.Run(&options.App{
		Title:  "Golden Fruit",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Linux: &linux.Options{
			Icon: appIcon,
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal("Error:", err.Error())
	}
}
