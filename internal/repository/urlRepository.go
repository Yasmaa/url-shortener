package repository

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/infrastructure/datastore"
)

type UrlRepository interface {
	// TODO
}

type urlRepository struct {
	db *datastore.Storage
}

func NewUrlRepository(db *datastore.Storage) UrlRepository {
	return &urlRepository{db: db}
}
