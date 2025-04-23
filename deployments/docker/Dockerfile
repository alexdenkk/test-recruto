FROM golang:1.24

LABEL maintaner="alexdenkk@yandex.ru"

WORKDIR app/

COPY . .

RUN go mod download

COPY go.mod go.sum ./

RUN go build -o recruto cmd/recruto/main.go
