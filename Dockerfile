FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o app

EXPOSE 3000

CMD ["./app"]