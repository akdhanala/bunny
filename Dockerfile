FROM golang:1.26.3-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bunny_app_binary .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/bunny_app_binary /bunny_app_binary

COPY --from=builder /app/config.yaml /config.yaml

EXPOSE 8080

ENTRYPOINT ["/bunny_app_binary"]