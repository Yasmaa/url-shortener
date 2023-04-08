package cmd

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/delivery/router"
	"github.com/vioma/url-shortener-api-pdceub/internal/infrastructure/datastore"
	"github.com/vioma/url-shortener-api-pdceub/internal/registry"
	"net/http"
)

func Start() {
	st := datastore.NewStorage()

	rg := registry.NewInteractor(st)

	h := rg.NewAppHandler()
	mux := router.NewRouter(h)

	http.ListenAndServe(":8000", mux)

}
