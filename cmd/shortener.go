package main

import (
	"github.com/ZaX51/url-shortener/internal/app"
)

func main() {
	app := app.New()

	app.Run(":3000")
}
