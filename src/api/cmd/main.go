package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/singfield/Horizon-chan/api/internal/models"
	"github.com/singfield/Horizon-chan/api/internal/services"
)

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func RunHslServeur() error{
	hslServeur := models.HslServer{"/songs", 8080}
	hslService := services.NewHslService(&hslServeur)
	err := hslService.Serve("/")
	if err !=nil {
		return err
	}
	return nil
}
func Run() error {
	app := fiber.New()
	err := app.Listen(":3000")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	go RunHslServeur()
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
