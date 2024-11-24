package usecase

import (
	"errors"
	"testing"

	"github.com/paulnune/goexpert-weather/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockWeatherService é um mock da interface WeatherService.
type MockWeatherService struct {
	mock.Mock
}

func (m *MockWeatherService) GetWeather(location string) (map[string]float64, error) {
	args := m.Called(location)
	if args.Get(0) != nil {
		return args.Get(0).(map[string]float64), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestWeatherUseCase_GetWeatherByZipCode(t *testing.T) {
	zipCodeRepo := &repository.MockZipCodeRepository{}
	weatherService := &MockWeatherService{}
	useCase := NewWeatherUseCase(zipCodeRepo, weatherService, "fake-key")

	t.Run("success", func(t *testing.T) {
		zipCodeRepo.On("GetLocationByZipCode", "01001000").Return(&repository.Location{
			City:         "São Paulo",
			Neighborhood: "Sé",
		}, nil)
		weatherService.On("GetWeather", "Sé, São Paulo").Return(map[string]float64{
			"temp_C": 25.0,
			"temp_F": 77.0,
			"temp_K": 298.15,
		}, nil)

		weather, err := useCase.GetWeatherByZipCode("01001000")

		assert.NoError(t, err)
		assert.Equal(t, 25.0, weather["temp_C"])
		assert.Equal(t, 77.0, weather["temp_F"])
		assert.Equal(t, 298.15, weather["temp_K"])

		zipCodeRepo.AssertExpectations(t)
		weatherService.AssertExpectations(t)
	})

	t.Run("zip code not found", func(t *testing.T) {
		zipCodeRepo.On("GetLocationByZipCode", "00000000").Return(nil, errors.New("zipcode not found"))

		_, err := useCase.GetWeatherByZipCode("00000000")

		assert.Error(t, err)
		assert.Equal(t, "zipcode not found", err.Error())

		zipCodeRepo.AssertExpectations(t)
	})

	t.Run("invalid zip code", func(t *testing.T) {
		zipCodeRepo.On("GetLocationByZipCode", "invalid").Return(nil, errors.New("invalid zipcode"))

		_, err := useCase.GetWeatherByZipCode("invalid")

		assert.Error(t, err)
		assert.Equal(t, "invalid zipcode", err.Error())

		zipCodeRepo.AssertExpectations(t)
	})
}
