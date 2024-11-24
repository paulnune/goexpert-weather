package repository

import "github.com/stretchr/testify/mock"

// MockZipCodeRepository implementa ZipCodeRepository
type MockZipCodeRepository struct {
	mock.Mock
}

// GetLocationByZipCode é o mock do método
func (m *MockZipCodeRepository) GetLocationByZipCode(zipCode string) (*Location, error) {
	args := m.Called(zipCode)
	if args.Get(0) != nil {
		return args.Get(0).(*Location), args.Error(1)
	}
	return nil, args.Error(1)
}
