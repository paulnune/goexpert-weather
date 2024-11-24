package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulnune/goexpert-weather/configs"
	"github.com/paulnune/goexpert-weather/internal/delivery/rest"
	"github.com/paulnune/goexpert-weather/internal/repository"
	"github.com/paulnune/goexpert-weather/internal/usecase"
)

func main() {
	// Carrega configurações
	configs.LoadConfig()

	// Inicializa repositórios e casos de uso
	zipCodeRepo := repository.NewZipCodeRepository()
	weatherUseCase := usecase.NewWeatherUseCase(zipCodeRepo, configs.GetConfig().WeatherAPIKey)

	// Configura roteador
	router := mux.NewRouter()
	handler := rest.NewHandler(weatherUseCase)

	// Rota base para verificar se o servidor está funcionando
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Servidor está funcionando"))
	}).Methods(http.MethodGet)

	// Rota para buscar clima por query string
	router.HandleFunc("/weather", handler.GetWeather).Methods(http.MethodGet)

	// Log de todas as rotas registradas
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		log.Printf("Rota registrada: %s\n", path)
		return nil
	})
	if err != nil {
		log.Fatalf("Erro ao listar rotas: %v\n", err)
	}

	// Inicializa o servidor
	log.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}
}
