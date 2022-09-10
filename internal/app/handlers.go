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

func (p *App) urlCut(c *fiber.Ctx) error {
	body := new(UrlCutRequest)

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": `BODY_ERROR`,
		})
	}

	err = p.validator.ValidateUrl(body.Url)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hash := encoder.Encode(body.Url, p.urlLength)

	err = p.url_storage.Set(hash, body.Url)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(UrlCutResponse{
		Url: fmt.Sprintf("%v/%v", p.domain, hash),
	})
}

func (p *App) urlOpen(c *fiber.Ctx) error {
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
