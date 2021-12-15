package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/singfield/Horizon-chan/api/internal/models"
)

type DefaultHslService struct {
	serveur models.HslServer
}

func NewHslService(hslServeur *models.HslServer) *DefaultHslService {
	return &DefaultHslService{
		serveur: *hslServeur,
	}
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func (dh *DefaultHslService) Serve(path string) error {

	http.Handle(path, addHeaders(http.FileServer(http.Dir(dh.serveur.SongsDir))))
	log.Printf("Starting Serveur on port %v\n", dh.serveur.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", dh.serveur.Port), nil)
	if err != nil {
		return err
	}
	return nil
}
