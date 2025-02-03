package product_app

import "net/http"

type App struct {
	url    string
	token  string
	client *http.Client
}

// TODO: add middleware

func New(url, token string, client *http.Client) *App {
	return &App{
		url:    url,
		token:  token,
		client: client,
	}
}
