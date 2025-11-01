package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/surajiyer26/go-url-shortener/handlers"
	"github.com/surajiyer26/go-url-shortener/storage"
)

func main() {
	store := storage.NewMemoryStore()

	http.HandleFunc("/shorten", handlers.ShortenHandler(store))
	http.HandleFunc("/", handlers.RedirectHandler(store))

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
