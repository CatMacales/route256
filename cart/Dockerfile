FROM golang:1.22.5-alpine AS builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /cart ./cmd/cart

FROM scratch
COPY --from=builder cart /bin/cart
COPY  config/prod.yaml /config.yaml

ENTRYPOINT ["/bin/cart"]
