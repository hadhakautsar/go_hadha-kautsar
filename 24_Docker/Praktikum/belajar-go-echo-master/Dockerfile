FROM golang:1.20.2

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 8080