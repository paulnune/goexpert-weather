package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// WeatherService define a interface para o serviço de clima
type WeatherService interface {
	GetWeather(location string) (map[string]float64, error)
}

// weatherAPIService é a implementação do WeatherService
type weatherAPIService struct {
	apiKey string
}

// NewWeatherService cria uma nova instância do WeatherService
func NewWeatherService(apiKey string) WeatherService {
	return &weatherAPIService{apiKey: apiKey}
}

// GetWeather busca o clima para uma localização específica
func (s *weatherAPIService) GetWeather(location string) (map[string]float64, error) {
	encodedLocation := url.QueryEscape(location)
	apiURL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", s.apiKey, encodedLocation)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, errors.New("não foi possível buscar os dados de clima")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta inesperada da WeatherAPI: %s", resp.Status)
	}

	var result struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.New("não foi possível decodificar os dados de clima")
	}

	return map[string]float64{
		"temp_C": result.Current.TempC,
		"temp_F": result.Current.TempC*1.8 + 32,
		"temp_K": result.Current.TempC + 273.15,
	}, nil
}
