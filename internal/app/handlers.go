package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZaX51/url-shortener/internal/encoder"
)

type UrlCutRequest struct {
	Url string `json:"url"`
}

type UrlCutResponse struct {
	Url string `json:"url"`
}

const urlLength = 7

func (p App) handleUrlCut(responseWriter http.ResponseWriter, request *http.Request) {
	var b UrlCutRequest

	err := json.NewDecoder(request.Body).Decode(&b)
	if err != nil {
		http.Error(responseWriter, "Invalid body", http.StatusBadRequest)
		return
	}

	if len(b.Url) == 0 {
		http.Error(responseWriter, `Missing or empty "url" field`, http.StatusBadRequest)
		return
	}

	hash := encoder.Encode(b.Url)[:urlLength]

	response := UrlCutResponse{
		Url: fmt.Sprintf("http://localhost:3000/%v", hash),
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(response)
}
