package services

import "github.com/stretchr/testify/mock"

type MockWeatherService struct {
	mock.Mock
}

func (m *MockWeatherService) GetWeather(location string) (map[string]float64, error) {
	args := m.Called(location)
	return args.Get(0).(map[string]float64), args.Error(1)
}
