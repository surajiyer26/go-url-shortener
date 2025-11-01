package handlers

import (
	"net/http"
	"strings"

	"github.com/surajiyer26/go-url-shortener/storage"
)

func RedirectHandler(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/")
		if id == "" {
			http.Error(w, "No short ID provided", http.StatusBadRequest)
			return
		}

		longURL, ok := store.Get(id)
		if !ok {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, longURL, http.StatusFound)
	}
}
