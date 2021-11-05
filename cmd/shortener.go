package main

import (
	"github.com/ZaX51/url-shortener/internal/app"
)

func main() {
	app := app.App{}

	app.Init(3000)
	app.Run()
}
