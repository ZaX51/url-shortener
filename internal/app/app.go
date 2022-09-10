package app

import (
	"log"
	"time"

	"github.com/ZaX51/url-shortener/internal/url_storage"
	"github.com/ZaX51/url-shortener/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Domain          string
	UrlLength       int
	RedisAddr       string
	RedisExpiration time.Duration
}

type App struct {
	domain      string
	urlLength   int
	server      *fiber.App
	url_storage *url_storage.UrlStorage
	validator   *validator.Validator
}

func New(config Config) *App {
	app := new(App)

	app.server = fiber.New(fiber.Config{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	})

	app.server.Post("/cut", app.urlCut)
	app.server.Get("/:hash", app.urlOpen)

	app.url_storage = &url_storage.UrlStorage{
		Addr:       config.RedisAddr,
		Expiration: config.RedisExpiration,
	}

	app.validator = validator.New(config.Domain)

	app.domain = config.Domain
	app.urlLength = config.UrlLength

	return app
}

func (p *App) Run(addr string) {
	err := p.url_storage.Connect()
	if err != nil {
		log.Fatal(err)
	}

	p.server.Listen(addr)
}
