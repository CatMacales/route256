LOCAL_BIN:=$(CURDIR)/bin
GLOBAL_BIN:=$(CURDIR)/../bin

.PHONY: run-cart
run-cart:
	CONFIG_PATH=./config/local.yaml go run ./cmd/cart

.PHONY: gen-mocks
gen-mocks:
	$(GLOBAL_BIN)/minimock -i github.com/CatMacales/route256/cart/internal/service/cart.CartProvider -o ./internal/service/cart/mock/cart_provider_mock.go -n CartProviderMock -p mock
	$(GLOBAL_BIN)/minimock -i github.com/CatMacales/route256/cart/internal/service/cart.ProductService -o ./internal/service/cart/mock/product_service_mock.go -n ProductServiceMock -p mock
	$(GLOBAL_BIN)/minimock -i github.com/CatMacales/route256/cart/internal/service/cart.LOMSService -o ./internal/service/cart/mock/loms_service_mock.go -n LOMSServiceMock -p mock
	go mod tidy

.PHONY: test-coverage
test-coverage:
	go test ./internal/service/cart ./internal/repository/memory/cart -coverprofile cover.out.tmp && \
	cat cover.out.tmp | grep -v "_mock.go" > cover.out && \
	rm cover.out.tmp && \
	go tool cover -func cover.out && \
	rm cover.out

.PHONY: test-integration
test-integration:
	go test ./test/integration -tags=integration,testmode

.PHONY: cyclo-report
cyclo-report:
	$(GLOBAL_BIN)/gocyclo -avg -top 10 -ignore ".*_mock.go|.*.pb.*" .

.PHONY: cognit-report
cognit-report:
	$(GLOBAL_BIN)/gocognit -avg -top 10 -ignore '.*_mock.go|.*.pb.*' .

get-loms-gen-proto:
	mkdir -p ./pkg/api
	cp -r ../loms/pkg/api/loms ./pkg/api