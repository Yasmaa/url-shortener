package usecase

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/repository"
)

type UrlService interface {
	// TODO : implement services
}

type urlService struct {
	UrlRepository repository.UrlRepository
}

func NewUrlService(ur repository.UrlRepository) UrlService {
	return &urlService{UrlRepository: ur}
}
