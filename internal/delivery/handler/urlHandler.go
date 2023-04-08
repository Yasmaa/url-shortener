package handler

import (
	"github.com/vioma/url-shortener-api-pdceub/internal/usecase"
	"github.com/vioma/url-shortener-api-pdceub/internal/domain"

	"net/http"
	"encoding/json"
)

type UrlHandler interface {

	Ping(w http.ResponseWriter, r *http.Request)
	Encode(w http.ResponseWriter, r *http.Request)
	Decode(w http.ResponseWriter, r *http.Request)



}

type urlHandler struct {
	UrlService usecase.UrlService
}

func NewUrlHandler(uc usecase.UrlService) UrlHandler {
	return &urlHandler{UrlService: uc}
}


func (uh *urlHandler) Ping(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("content-type", "application/json")
		message, _ := json.Marshal(map[string]string{"message": "pong"})
		w.Write(message)

	
}

func (uh *urlHandler) Encode(w http.ResponseWriter, r *http.Request) {

	var url domain.URLRequest

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		
		w.WriteHeader(http.StatusBadRequest)
		message, _ := json.Marshal(map[string]string{"message": "bad request"})
		w.Write(message)
		return
	}

	res, err := uh.UrlService.Encode(url.LongUrl)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		message, _ := json.Marshal(map[string]string{"message": "internal server error"})
		w.Write(message)
		return

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)






	
}

func (uh *urlHandler) Decode(w http.ResponseWriter, r *http.Request) {

	

	
}