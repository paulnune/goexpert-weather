package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/paulnune/goexpert-weather/internal/repository"
)

// ZipCodeRepository define a interface para o repositório de CEP
type ZipCodeRepository interface {
	GetLocationByZipCode(zipCode string) (*repository.Location, error)
}

type WeatherUseCase struct {
	zipCodeRepo ZipCodeRepository
	apiKey      string
}

func NewWeatherUseCase(zipCodeRepo ZipCodeRepository, apiKey string) *WeatherUseCase {
	return &WeatherUseCase{
		zipCodeRepo: zipCodeRepo,
		apiKey:      apiKey,
	}
}

func (u *WeatherUseCase) GetWeatherByZipCode(zipCode string) (map[string]float64, error) {
	// Obtém a localização usando o repositório
	location, err := u.zipCodeRepo.GetLocationByZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	// Se o bairro ou logradouro estiverem disponíveis, inclua na consulta
	var query string
	if location.Neighborhood != "" && location.City != "" {
		query = fmt.Sprintf("%s, %s", location.Neighborhood, location.City)
	} else if location.City != "" {
		query = location.City
	} else {
		return nil, errors.New("could not determine location for the given zip code")
	}

	// Substituir espaços no nome da localidade por `%20`
	formattedQuery := url.QueryEscape(query)
	weatherAPIURL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", u.apiKey, formattedQuery)

	log.Printf("URL para consulta WeatherAPI: %s\n", weatherAPIURL)

	resp, err := http.Get(weatherAPIURL)
	if err != nil {
		return nil, errors.New("could not fetch weather data")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Resposta inesperada da WeatherAPI. Status: %d, Body: %s\n", resp.StatusCode, string(body))
		return nil, errors.New("could not fetch weather data")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("could not read weather data")
	}

	var weatherData struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, errors.New("could not parse weather data")
	}

	return map[string]float64{
		"temp_C": weatherData.Current.TempC,
		"temp_F": weatherData.Current.TempC*1.8 + 32,
		"temp_K": weatherData.Current.TempC + 273.15,
	}, nil
}
