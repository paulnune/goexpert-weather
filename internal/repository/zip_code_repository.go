package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ZipCodeRepository define a interface para o repositório de CEPs
type ZipCodeRepository interface {
	GetLocationByZipCode(zipCode string) (*Location, error)
}

// zipCodeRepository é a implementação concreta da interface ZipCodeRepository
type zipCodeRepository struct{}

// NewZipCodeRepository cria uma nova instância de zipCodeRepository
func NewZipCodeRepository() ZipCodeRepository {
	return &zipCodeRepository{}
}

// Location representa a estrutura de dados retornada pela API ViaCEP
type Location struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// GetLocationByZipCode busca a localização com base no CEP fornecido
func (r *zipCodeRepository) GetLocationByZipCode(zipCode string) (*Location, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("erro ao buscar o CEP")
	}

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	if location.Cep == "" {
		return nil, errors.New("CEP não encontrado")
	}

	return &location, nil
}
