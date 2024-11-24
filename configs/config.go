package configs

import (
	"log"
	"os"
)

type Config struct {
	WeatherAPIKey string
}

var config *Config

func LoadConfig() {
	config = &Config{
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
	}

	if config.WeatherAPIKey == "" {
		log.Fatal("Chave da API do WeatherAPI não configurada (WEATHER_API_KEY)")
	}

	log.Println("Configurações carregadas com sucesso")
}

func GetConfig() *Config {
	if config == nil {
		log.Fatal("Configurações não carregadas. Certifique-se de chamar LoadConfig antes de GetConfig.")
	}
	return config
}
