FROM golang:1.23-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
RUN go build -o app ./cmd/app

CMD ["./app"]
