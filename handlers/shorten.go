package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/surajiyer26/go-url-shortener/storage"
)

type request struct {
	URL string `json:"url"`
}

func generateShortID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

func ShortenHandler(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id := generateShortID()
		store.Save(id, req.URL)

		resp := map[string]string{"short_url": "http://localhost:8080/" + id}
		json.NewEncoder(w).Encode(resp)
	}
}
