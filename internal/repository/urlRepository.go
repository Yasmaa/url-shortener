package repository

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/domain"
	"github.com/vioma/url-shortener-api-pdceub/internal/infrastructure/datastore"

	"github.com/vioma/url-shortener-api-pdceub/utils/error_utils"
)

type UrlRepository interface {
	Get(alias string) (string, error_utils.MessageErr)
	Exists(alias string) (bool, error_utils.MessageErr)
	Store(url, alias string) (*domain.URL, error_utils.MessageErr)
}

type urlRepository struct {
	db *datastore.Storage
}

func NewUrlRepository(db *datastore.Storage) UrlRepository {
	return &urlRepository{db: db}
}

func (ur *urlRepository) Get(alias string) (string, error_utils.MessageErr) {
	url, found := ur.db.Urls[alias]
	if !found {
		return "", error_utils.NewNotFoundError("not found")
	}

	return url, nil
}

// Exists checks if there is a URL with the requested alias
func (ur *urlRepository) Exists(alias string) (bool, error_utils.MessageErr) {
	_, found := ur.db.Urls[alias]
	if found {
		return found,error_utils.NewBadRequestError("already  exists")
	}
	return found, nil
}

// Store creates a new short URL
func (ur *urlRepository) Store(url, alias string) (*domain.URL, error_utils.MessageErr) {
	
	ur.db.Urls[alias] = url

	res := &domain.URL{
		LongUrl: url,
		Alias:   alias}
	return res, nil
}
