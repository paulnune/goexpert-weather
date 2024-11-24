package usecase

import (
	"errors"
	"fmt"

	"github.com/paulnune/goexpert-weather/internal/repository"
)

// WeatherService define a interface para o serviço de clima
type WeatherService interface {
	GetWeather(location string) (map[string]float64, error)
}

// WeatherUseCase define o caso de uso
type WeatherUseCase struct {
	zipCodeRepo    repository.ZipCodeRepository
	weatherService WeatherService
	apiKey         string
}

// NewWeatherUseCase cria uma nova instância do caso de uso
func NewWeatherUseCase(zipCodeRepo repository.ZipCodeRepository, weatherService WeatherService, apiKey string) *WeatherUseCase {
	return &WeatherUseCase{
		zipCodeRepo:    zipCodeRepo,
		weatherService: weatherService,
		apiKey:         apiKey,
	}
}

func (u *WeatherUseCase) GetWeatherByZipCode(zipCode string) (map[string]float64, error) {
	location, err := u.zipCodeRepo.GetLocationByZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	var query string
	if location.Neighborhood != "" && location.City != "" {
		query = fmt.Sprintf("%s, %s", location.Neighborhood, location.City)
	} else if location.City != "" {
		query = location.City
	} else {
		return nil, errors.New("could not determine location for the given zip code")
	}

	return u.weatherService.GetWeather(query)
}
