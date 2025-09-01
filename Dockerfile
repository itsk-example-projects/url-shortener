FROM golang:1.25-alpine AS builder

WORKDIR /build

RUN apk add --no-cache upx

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -v -ldflags="-s -w" -o url-shortener-example cmd/main.go

RUN upx --best --lzma url-shortener-example

FROM scratch

COPY --from=builder /build/url-shortener-example ./
COPY --from=builder /build/internal/web/templates ./internal/web/templates
COPY --from=builder /build/config.yaml ./

EXPOSE 8080

CMD ["./url-shortener-example"]
