package router

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/delivery/handler"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, handler handler.AppHandler) {
	// TODO : implement routes

}

func NewRouter(handler handler.AppHandler) *http.ServeMux {

	mux := http.NewServeMux()
	SetupRoutes(mux, handler)

	return mux
}
