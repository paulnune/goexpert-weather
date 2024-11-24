package usecase

import (
	"errors"
	"testing"

	"github.com/paulnune/goexpert-weather/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock para o ZipCodeRepository
type MockZipCodeRepository struct{}

func (m *MockZipCodeRepository) GetLocationByZipCode(zipCode string) (*repository.Location, error) {
	if zipCode == "01001000" {
		return &repository.Location{City: "São Paulo", Neighborhood: "Sé"}, nil
	}
	if zipCode == "00000000" {
		return nil, errors.New("zipcode not found")
	}
	return nil, errors.New("invalid zipcode")
}

// Mock para simular a API de clima
type MockWeatherService struct{}

func (m *MockWeatherService) GetWeather(city string) (map[string]float64, error) {
	if city == "Sé, São Paulo" {
		return map[string]float64{"temp_C": 25, "temp_F": 77, "temp_K": 298.15}, nil
	}
	return nil, errors.New("could not fetch weather data")
}

func TestWeatherUseCase_GetWeatherByZipCode(t *testing.T) {
	zipCodeRepo := &MockZipCodeRepository{}
	useCase := NewWeatherUseCase(zipCodeRepo, "fake-key")

	t.Run("success", func(t *testing.T) {
		weather, err := useCase.GetWeatherByZipCode("01001000")
		require.NoError(t, err)
		require.NotNil(t, weather)
		assert.Equal(t, 25.0, weather["temp_C"])
		assert.Equal(t, 77.0, weather["temp_F"])
		assert.Equal(t, 298.15, weather["temp_K"])
	})

	t.Run("zip code not found", func(t *testing.T) {
		_, err := useCase.GetWeatherByZipCode("00000000")
		require.Error(t, err)
		assert.Equal(t, "zipcode not found", err.Error())
	})

	t.Run("invalid zip code", func(t *testing.T) {
		_, err := useCase.GetWeatherByZipCode("invalid")
		require.Error(t, err)
		assert.Equal(t, "invalid zipcode", err.Error())
	})
}
