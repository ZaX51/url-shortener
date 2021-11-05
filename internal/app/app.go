package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const apiRoutesPrefix = "/api"

type App struct {
	server *http.Server
}

func (p *App) Init(port int) {
	r := mux.NewRouter()
	s := r.PathPrefix(apiRoutesPrefix).Subrouter()

	s.HandleFunc("/cut", p.handleUrlCut).Methods(http.MethodPost)

	p.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      s,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
}

func (p *App) Run() {
	fmt.Printf("Server listening on %s\n", p.server.Addr)
	log.Fatal(p.server.ListenAndServe())
}
