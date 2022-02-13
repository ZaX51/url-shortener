package app

import (
	"log"
	"time"

	"github.com/ZaX51/url-shortener/internal/url_storage"
	"github.com/ZaX51/url-shortener/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	server      *fiber.App
	url_storage *url_storage.UrlStorage
	validator   *validator.Validator
}

func New() *App {
	app := new(App)

	app.server = fiber.New(fiber.Config{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	})

	app.server.Post("/cut", app.handleUrlCut)
	app.server.Get("/:hash", app.handleUrlOpen)

	app.url_storage = &url_storage.UrlStorage{
		Addr:       "localhost:6379",
		Expiration: 20 * time.Second,
	}

	app.validator = validator.New()

	return app
}

func (p *App) Run(addr string) {
	err := p.url_storage.Connect()
	if err != nil {
		log.Fatal(err)
	}

	p.server.Listen(addr)
}
