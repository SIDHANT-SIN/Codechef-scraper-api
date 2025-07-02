FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -ldflags "-s -w" -o /codechef-scraper-api ./main.go

FROM alpine:latest

COPY --from=builder /codechef-scraper-api /usr/local/bin/codechef-scraper-api

ENTRYPOINT ["/usr/local/bin/codechef-scraper-api"]