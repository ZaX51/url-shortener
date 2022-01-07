package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ZaX51/url-shortener/internal/url_storage"
	"github.com/gorilla/mux"
)

type App struct {
	server      *http.Server
	url_storage *url_storage.UrlStorage
}

func (p *App) Init(port int) {
	r := mux.NewRouter()

	r.HandleFunc("/cut", p.handleUrlCut).Methods(http.MethodPost)
	r.HandleFunc("/{hash}", p.handleUrlOpen).Methods(http.MethodGet)

	p.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	p.url_storage = &url_storage.UrlStorage{
		Addr:       "localhost:6379",
		Expiration: 20 * time.Second,
	}
}

func (p *App) Run() {
	err := p.url_storage.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server listening on %s\n", p.server.Addr)
	log.Fatal(p.server.ListenAndServe())
}
