FROM golang:1.22.5-alpine AS builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /loms ./cmd/loms

FROM scratch
COPY --from=builder loms /bin/loms
COPY  config/prod.yaml /config.yaml
COPY  stock-data.json /stock-data.json

ENTRYPOINT ["/bin/loms"]
