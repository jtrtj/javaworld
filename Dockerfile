FROM golang:1.15.2-buster

WORKDIR /go/src/app

COPY go.mod .
COPY . .
RUN go mod download

CMD go run main.go
