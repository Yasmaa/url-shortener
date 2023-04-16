package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/vioma/url-shortener-api-pdceub/internal/delivery/router"
	"github.com/vioma/url-shortener-api-pdceub/internal/infrastructure/datastore"
	"github.com/vioma/url-shortener-api-pdceub/internal/registry"
)


func init() {
	//loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Printf("sad .env file found")
	}
}


func Start() {
	st := datastore.NewStorage()

	rg := registry.NewInteractor(st)

	h := rg.NewAppHandler()
	mux := router.NewRouter(h)

	log.Printf("HTTP Server listening at  %s", fmt.Sprintf(":%s", os.Getenv("PORT")))

	httpServer := http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: mux,
	}

	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(idleConnectionsClosed)
	}()

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	} 


	<-idleConnectionsClosed

	log.Printf("Bye bye")

}
