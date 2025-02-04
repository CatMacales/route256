package product_app

import (
	"fmt"
	"net/http"
)

type App struct {
	url    string
	token  string
	client *http.Client
}

func New(url, token string, client *http.Client) *App {
	return &App{
		url:    url,
		token:  token,
		client: client,
	}
}

func handleHTTPError(statusCode int) error {
	switch statusCode {
	case http.StatusNotFound:
		return ErrProductNotFound
	case http.StatusUnauthorized:
		return ErrInvalidToken
	default:
		return fmt.Errorf("unexpected status code: %d", statusCode)
	}
}
