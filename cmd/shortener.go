package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ZaX51/url-shortener/internal/encoder"
	"github.com/go-redis/redis/v8"
)

var (
	ctx context.Context
	rdb *redis.Client
)

func main() {
	port := ":3000"
	ctx = context.TODO()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/cut", cutListener)
	http.HandleFunc("/get", getListener)

	fmt.Printf("Server listening on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func cutListener(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	url := request.Form.Get("url")
	hash := encoder.Encode(url)[:7]

	rdb.SetEX(ctx, hash, url, 30*time.Minute)

	hashedUrl := fmt.Sprintf("http://localhost:3000/get?hash=%v", hash)

	fmt.Fprintf(responseWriter, "<a href='%s'>%s</a>\n", hashedUrl, hashedUrl)
}

func getListener(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	hash := request.Form.Get("hash")

	if len(hash) == 0 {
		http.NotFound(responseWriter, request)
		return
	}

	url, err := rdb.Get(ctx, hash).Result()
	if err != nil {
		fmt.Printf("Hash not found: %s\n", hash)
		http.NotFound(responseWriter, request)
		return
	}

	http.Redirect(responseWriter, request, url, http.StatusFound)
}
