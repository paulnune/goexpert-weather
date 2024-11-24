package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulnune/goexpert-weather/configs"
	"github.com/paulnune/goexpert-weather/internal/delivery/rest"
	"github.com/paulnune/goexpert-weather/internal/repository"
	"github.com/paulnune/goexpert-weather/internal/services"
	"github.com/paulnune/goexpert-weather/internal/usecase"
)

func main() {
	// Carrega configurações
	configs.LoadConfig()

	// Inicializa repositório e serviço
	zipCodeRepo := repository.NewZipCodeRepository()
	weatherService := services.NewWeatherService() // Ajuste para criar uma implementação real

	// Inicializa caso de uso com o repositório, serviço e chave de API
	handler := rest.NewHandler(usecase.NewWeatherUseCase(zipCodeRepo, weatherService, configs.APIKey))

	// Configura roteador
	router := mux.NewRouter()

	// Define rotas
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Servidor está funcionando"))
	}).Methods(http.MethodGet)

	router.HandleFunc("/weather", handler.GetWeather).Methods(http.MethodGet)

	// Inicia o servidor
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
