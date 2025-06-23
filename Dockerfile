FROM golang:1.23-alpine

WORKDIR /app

# RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -a -o ./bin/api ./cmd/app

CMD ["/app/bin/api"]

EXPOSE 9999