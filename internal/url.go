package internal

import (
	"time"

	"github.com/google/uuid"
)

var urlHash = make(map[string]string)

type Url struct {
	ID           int
	uuid         string
	original_url string
	created_at   time.Time
	expires_at   *time.Time
}

func ShortenUrl(urlEndpoint string, expiration ...time.Time) Url {
	var expirationDate *time.Time // Use a pointer to handle optional expiration

	// Check if an expiration date was provided
	if len(expiration) > 0 {
		expirationDate = &expiration[0] // Set the expiration date
	}

	url := Url{
		ID:           1, // Must get the last ID
		uuid:         uuid.New().String(),
		original_url: urlEndpoint,
		created_at:   time.Now().Truncate(time.Second), // Truncate time to be like TIMESTAMP in MySQL
		expires_at:   expirationDate,
	}

	urlHash[url.uuid] = url.original_url

	return url
}
