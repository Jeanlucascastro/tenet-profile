package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *AuthClient) ValidateToken(token string) (bool, error) {
	log.Println("ValidateToken")

	if token == "" {
		return false, fmt.Errorf("token vazio")
	}

	token = strings.TrimPrefix(token, "Bearer ")

	print("sending request")
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/validate", c.baseURL), nil)
	print("Request ")
	print(req)
	print(err)
	if err != nil {
		return false, err
	}

	req.Header.Set("token", token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("token inválido: status %d", resp.StatusCode)
	}

	var valid bool
	if err := json.NewDecoder(resp.Body).Decode(&valid); err != nil {
		return false, err
	}

	return valid, nil
}
