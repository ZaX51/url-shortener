package app

import (
	"fmt"
	"net/http"

	"github.com/ZaX51/url-shortener/internal/encoder"
	"github.com/gofiber/fiber/v2"
)

type UrlCutRequest struct {
	Url string `json:"url"`
}

type UrlCutResponse struct {
	Url string `json:"url"`
}

const urlLength = 7

func (p *App) handleUrlCut(c *fiber.Ctx) error {
	body := new(UrlCutRequest)

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": `Cannot parse JSON`,
		})
	}

	err = p.validator.ValidateUrl(body.Url)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hash := encoder.Encode(body.Url, urlLength)

	err = p.url_storage.Set(hash, body.Url)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(UrlCutResponse{
		Url: fmt.Sprintf("http://localhost:3000/%v", hash),
	})
}

func (p *App) handleUrlOpen(c *fiber.Ctx) error {
	hash := c.Params("hash")

	url, err := p.url_storage.Get(hash)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	if len(url) == 0 {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.Redirect(url, http.StatusSeeOther)
}
