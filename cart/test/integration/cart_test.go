//go:build integration

package integration

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/app/loms"
	"github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/app/server"
	"github.com/CatMacales/route256/cart/internal/http-server/handler/cart"
	"github.com/CatMacales/route256/cart/internal/http-server/middleware"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"github.com/CatMacales/route256/cart/internal/repository/memory/cart"
	"github.com/CatMacales/route256/cart/internal/service/cart"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegration_Run(t *testing.T) {
	suite.Run(t, new(CartSuite))
}

type CartSuite struct {
	suite.Suite
	server         *httptest.Server
	cartRepo       *cart_repository.Repository
	productService *product_app.App
}

type Config struct {
	ProductService ProductServiceConfig
	LOMSService    LOMSServiceConfig
}

type ProductServiceConfig struct {
	URL   string
	Token string
}

type LOMSServiceConfig struct {
	URL string
}

func (cs *CartSuite) SetupSuite() {
	cfg := Config{
		ProductService: ProductServiceConfig{
			URL:   "http://route256.pavl.uk:8080",
			Token: "testtoken",
		},
		LOMSService: LOMSServiceConfig{
			URL: "localhost:50050",
		},
	}

	productApp := product_app.New(cfg.ProductService.URL, cfg.ProductService.Token, &http.Client{Transport: middleware.NewRetry(http.DefaultTransport)})
	cs.productService = productApp

	lomsApp := loms_app.New(cfg.LOMSService.URL)

	cartRepository := cart_repository.NewRepository()
	cs.cartRepo = cartRepository

	cartService := cart.NewService(cartRepository, productApp, lomsApp)

	handler := cart_handler.New(cartService)

	validation.InitValidator() // init beauty validator

	cs.server = httptest.NewServer(server.New("", 0, handler).Router)
}

func (cs *CartSuite) TearDownSuite() {
	cs.server.Close()
}

func (cs *CartSuite) TearDownTest() {
	cs.cartRepo.Clear(context.Background())
}
