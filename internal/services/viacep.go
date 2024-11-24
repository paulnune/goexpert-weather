package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Location struct {
	Localidade string `json:"localidade"`
}

func GetLocation(zipcode string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch location")
	}

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return "", err
	}

	return location.Localidade, nil
}
