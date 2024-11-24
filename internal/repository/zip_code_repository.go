package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ZipCodeRepository é a implementação da interface para buscar dados do CEP.
type ZipCodeRepository struct{}

// NewZipCodeRepository cria uma nova instância de ZipCodeRepository.
func NewZipCodeRepository() *ZipCodeRepository {
	return &ZipCodeRepository{}
}

// Location representa os dados retornados pela API ViaCEP.
type Location struct {
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
}

// GetLocationByZipCode busca informações de localização pelo CEP.
func (r *ZipCodeRepository) GetLocationByZipCode(zipCode string) (*Location, error) {
	// URL da API ViaCEP
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)

	// Faz a requisição
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("não foi possível buscar os dados do CEP: %w", err)
	}
	defer resp.Body.Close()

	// Verifica o status HTTP da resposta
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("CEP não encontrado")
	}

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("não foi possível ler os dados do CEP: %w", err)
	}

	// Estrutura para processar os dados retornados pela API ViaCEP
	var data struct {
		Localidade string `json:"localidade"`
		Bairro     string `json:"bairro"`
		Erro       bool   `json:"erro"`
	}

	// Faz o unmarshal do JSON
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("não foi possível processar os dados do CEP: %w", err)
	}

	// Verifica se a API retornou erro
	if data.Erro {
		return nil, errors.New("CEP não encontrado")
	}

	// Retorna os dados de localização
	return &Location{
		City:         data.Localidade,
		Neighborhood: data.Bairro,
	}, nil
}
