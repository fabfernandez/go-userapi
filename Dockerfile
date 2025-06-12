FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

# Install curl for healthcheck
RUN apk --no-cache add curl

WORKDIR /app

COPY --from=builder /app/main .
COPY docs/ ./docs/

EXPOSE 8080

CMD ["./main"] 