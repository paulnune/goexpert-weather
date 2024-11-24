package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/paulnune/goexpert-weather/internal/usecase"
)

type Handler struct {
	usecase *usecase.WeatherUseCase
}

func NewHandler(u *usecase.WeatherUseCase) *Handler {
	return &Handler{
		usecase: u,
	}
}

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	weather, err := h.usecase.GetWeatherByZipCode(cep)
	if err != nil {
		log.Printf("Erro ao buscar clima: %v", err)
		if err.Error() == "zipcode not found" {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
