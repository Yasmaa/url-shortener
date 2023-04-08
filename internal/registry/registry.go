package registry

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/delivery/handler"
	"github.com/vioma/url-shortener-api-pdceub/internal/infrastructure/datastore"
	"github.com/vioma/url-shortener-api-pdceub/internal/repository"
	"github.com/vioma/url-shortener-api-pdceub/internal/usecase"
)

type interactor struct {
	db *datastore.Storage
}

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(ds *datastore.Storage) Interactor {
	return &interactor{db: ds}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return handler.AppHandler{
		UrlHandler: i.NewUrlHandler()}
}

// URL
func (i *interactor) NewUrlHandler() handler.UrlHandler {
	return handler.NewUrlHandler(i.NewUrlService())
}

func (i *interactor) NewUrlService() usecase.UrlService {
	return usecase.NewUrlService(i.NewUrlRepository())
}

func (i *interactor) NewUrlRepository() repository.UrlRepository {
	return repository.NewUrlRepository(i.db)
}
