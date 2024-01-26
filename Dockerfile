FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o productapi product-api/cmd/main.go

CMD ["./productapi"]