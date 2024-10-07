FROM golang:1.23.2-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod vendor

RUN go build -o ./bin/server cmd/main.go

FROM alpine AS runner

COPY --from=builder /build/bin/server /
COPY config/config.yml /config/

ENTRYPOINT ["/server"]

#