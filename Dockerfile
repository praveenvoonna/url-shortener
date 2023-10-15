FROM golang:latest

WORKDIR /url-shortener

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o url-shortener .

EXPOSE 8080

CMD ["./url-shortener"]
