package storage

import "github.com/shortener/internal/app/models"

type Storage struct {
	URL map[string]models.URL
}

func NewStorage() Storage {
	return Storage{
		URL: make(map[string]models.URL),
	}
}
