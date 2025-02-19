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

vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis && \
 	cd vendor-proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/grpc-ecosystem && \
 	cd vendor-proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/grpc-ecosystem

.PHONY: vendor-rm
vendor-rm:
	rm -rf vendor-proto

vendor-proto: vendor-rm vendor-proto/validate vendor-proto/google/api vendor-proto/protoc-gen-openapiv2/options

.PHONY: bin-deps
bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5

LOMS_PROTO_PATH:=api/loms/v1

.PHONY: protoc-generate
protoc-generate: bin-deps vendor-proto
	mkdir -p api/openapiv2
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
        --plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
        --grpc-gateway_out pkg/${LOMS_PROTO_PATH} \
        --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
        --plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
        --openapiv2_out api/openapiv2 \
        --openapiv2_opt logtostderr=true \
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

.PHONY: serve-swagger
serve-swagger:
	bin/swagger serve -F=swagger api/openapiv2/loms.swagger.json