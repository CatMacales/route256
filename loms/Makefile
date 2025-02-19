LOCAL_BIN:=$(CURDIR)/bin
GLOBAL_BIN:=$(CURDIR)/../bin

PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor-proto/validate
		mv vendor-proto/tmp/validate vendor-proto/
		rm -rf vendor-proto/tmp

.PHONY: vendor-rm
vendor-rm:
	rm -rf vendor-proto

vendor-proto: vendor-rm vendor-proto/validate

.PHONY: bin-deps
bin-deps:
	$(info Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4

LOMS_PROTO_PATH:=api/loms/v1

.PHONY: protoc-generate
protoc-generate: bin-deps vendor-proto
	protoc \
		-I ${LOMS_PROTO_PATH} \
		-I vendor-proto \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
		--go_out pkg/${LOMS_PROTO_PATH} \
		--go_opt paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
		--go-grpc_out pkg/${LOMS_PROTO_PATH} \
		--go-grpc_opt paths=source_relative \
		--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
        --validate_out="lang=go,paths=source_relative:pkg/${LOMS_PROTO_PATH}" \
		${LOMS_PROTO_PATH}/loms.proto
	go mod tidy

.PHONY: gen-mocks
gen-mocks:
	$(GLOBAL_BIN)/minimock -i github.com/CatMacales/route256/loms/internal/service/loms.OrderProvider -o ./internal/service/loms/mock/order_provider_mock.go -n OrderProviderMock -p mock
	$(GLOBAL_BIN)/minimock -i github.com/CatMacales/route256/loms/internal/service/loms.StockProvider -o ./internal/service/loms/mock/stock_provider_mock.go -n StockProviderMock -p mock
	go mod tidy

.PHONY: test-coverage
test-coverage:
	go test ./internal/service/loms -coverprofile cover.out.tmp && cat cover.out.tmp | grep -v "_mock.go" > cover.out && rm cover.out.tmp && go tool cover -func cover.out && rm cover.out


.PHONY: cyclo-report
cyclo-report:
	$(GLOBAL_BIN)/gocyclo -avg -top 10 -ignore ".*_mock.go|.*.pb.*" .

.PHONY: cognit-report
cognit-report:
	$(GLOBAL_BIN)/gocognit -avg -top 10 -ignore '.*_mock.go|.*.pb.*' .