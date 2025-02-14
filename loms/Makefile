LOCAL_BIN:=$(CURDIR)/bin

PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

.PHONY: bin-deps
bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0



LOMS_PROTO_PATH:=api/loms/v1

.PHONY: protoc-generate
protoc-generate: bin-deps
	protoc \
		-I ${LOMS_PROTO_PATH} \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
		--go_out pkg/${LOMS_PROTO_PATH} \
		--go_opt paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
		--go-grpc_out pkg/${LOMS_PROTO_PATH} \
		--go-grpc_opt paths=source_relative \
		${LOMS_PROTO_PATH}/loms.proto
	go mod tidy
