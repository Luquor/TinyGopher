package internal

import (
	"errors"
	"log"
	"net/http"
	"time"
)

var URLList []URL

type URL struct {
	ID           int
	uuid         string
	original_url string
	created_at   time.Time
	expires_at   time.Time
}

func findOrinalURL(uuid string) (string, error) {
	for i := 0; i < len(URLList); i++ {
		if uuid == URLList[i].uuid {
			return URLList[i].original_url, nil
		}
	}
	return "", errors.New("the UUID does not exist")
}

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Path[5:]

	if uuid == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You must provide an uuid"))
	}

	originalURL, err := findOrinalURL(uuid)
	if err != nil {
		log.Fatalf("Server failed to get UUID: %v", err)
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
