package client

import "net/http"

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

}
