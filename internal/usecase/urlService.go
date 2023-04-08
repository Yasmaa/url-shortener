package usecase

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/repository"
	"github.com/vioma/url-shortener-api-pdceub/internal/domain"

	"github.com/vioma/url-shortener-api-pdceub/utils/error_utils"
	"github.com/vioma/url-shortener-api-pdceub/utils/alias"


)

type UrlService interface {
	Encode(url string) (*domain.Url, error_utils.MessageErr)
	Decode(alias string) (string, error_utils.MessageErr)
}

type urlService struct {
	UrlRepository repository.UrlRepository
}

func NewUrlService(ur repository.UrlRepository) UrlService {
	return &urlService{UrlRepository: ur}
}


func (us *urlService) Encode(longUrl string) (*domain.Url, error_utils.MessageErr) {
	al := alias.GenerateBase62ID(8)
	return us.UrlRepository.Store(longUrl, al)

}

func (us *urlService) Decode(alias string) (string, error_utils.MessageErr) {
	return us.UrlRepository.Get(alias)

}
