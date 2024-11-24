package usecase

import (
	"errors"
	"fmt"

	"github.com/paulnune/goexpert-weather/internal/repository"
	"github.com/paulnune/goexpert-weather/internal/services"
)

// WeatherUseCase define o caso de uso para obter o clima por CEP
type WeatherUseCase struct {
	zipCodeRepo    repository.ZipCodeRepository
	weatherService services.WeatherService
}

// NewWeatherUseCase cria uma nova instância de WeatherUseCase
func NewWeatherUseCase(zipCodeRepo repository.ZipCodeRepository, weatherService services.WeatherService) *WeatherUseCase {
	return &WeatherUseCase{
		zipCodeRepo:    zipCodeRepo,
		weatherService: weatherService,
	}
}

// GetWeatherByZipCode obtém o clima com base no CEP fornecido
func (u *WeatherUseCase) GetWeatherByZipCode(zipCode string) (map[string]float64, error) {
	location, err := u.zipCodeRepo.GetLocationByZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	var query string
	if location.Bairro != "" && location.Localidade != "" {
		query = fmt.Sprintf("%s, %s", location.Bairro, location.Localidade)
	} else if location.Localidade != "" {
		query = location.Localidade
	} else {
		return nil, errors.New("não foi possível determinar a localização para o CEP fornecido")
	}

	return u.weatherService.GetWeather(query)
}
