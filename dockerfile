FROM golang:1.24.0

WORKDIR /rest

COPY . /rest

RUN go mod download

RUN go build -o run ./cmd/main.go

CMD ["./run"]
