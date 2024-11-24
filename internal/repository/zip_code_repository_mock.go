package repository

import "github.com/stretchr/testify/mock"

type MockZipCodeRepository struct {
	mock.Mock
}

func (m *MockZipCodeRepository) GetLocationByZipCode(zipCode string) (*Location, error) {
	args := m.Called(zipCode)
	if args.Get(0) != nil {
		return args.Get(0).(*Location), args.Error(1)
	}
	return nil, args.Error(1)
}
