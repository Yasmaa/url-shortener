package handler

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/usecase"
)

type UrlHandler interface {

	// TODO : implement handlers
}

type urlHandler struct {
	UrlService usecase.UrlService
}

func NewUrlHandler(uc usecase.UrlService) UrlHandler {
	return &urlHandler{UrlService: uc}
}
