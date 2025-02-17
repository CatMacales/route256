CART_BIN:=$(CURDIR)/bin

.PHONY: run-cart
run-cart:
	CONFIG_PATH=./config/local.yaml go run ./cmd/cart

.PHONY: bin-deps
bin-deps:
	GOBIN=$(CART_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.4
	GOBIN=$(CART_BIN) go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.6.0
	GOBIN=$(CART_BIN) go install github.com/uudashr/gocognit/cmd/gocognit@v1.2.0

.PHONY: gen-mocks
gen-mocks:
	$(CART_BIN)/minimock -i github.com/CatMacales/route256/cart/internal/service/cart.CartRepository -o ./internal/service/cart/mock/cart_repository_mock.go -n CartRepositoryMock -p mock
	$(CART_BIN)/minimock -i github.com/CatMacales/route256/cart/internal/service/cart.ProductService -o ./internal/service/cart/mock/product_service_mock.go -n ProductServiceMock -p mock
	go mod tidy

.PHONY: test-coverage
test-coverage:
	go test ./internal/service/cart ./internal/repository/memory/cart -coverprofile cover.out.tmp && cat cover.out.tmp | grep -v "_mock.go" > cover.out && rm cover.out.tmp && go tool cover -func cover.out && rm cover.out

.PHONY: cyclo-report
cyclo-report:
	$(CART_BIN)/gocyclo -avg -top 10 -ignore ".*_mock.go" .

cognit-report:
	$(CART_BIN)/gocognit -avg -top 10 -ignore '.*_mock.go' .

get-loms-gen-proto:
	mkdir -p ./pkg/api
	cp -r ../loms/pkg/api/loms ./pkg/api