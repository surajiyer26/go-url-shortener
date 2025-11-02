package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/surajiyer26/go-url-shortener/handlers"
	"github.com/surajiyer26/go-url-shortener/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	redisAddr := os.Getenv("REDIS_ADDR")

	var store storage.Store
	if redisAddr != "" {
		store = storage.NewRedisStore(redisAddr)
	} else {
		store = storage.NewMemoryStore()
	}

	http.HandleFunc("/shorten", handlers.ShortenHandler(store))
	http.HandleFunc("/", handlers.RedirectHandler(store))

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
