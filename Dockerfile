FROM golang:1.22-bullseye

RUN apt update

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Change .env with docker
COPY .env.docker ./.env

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]