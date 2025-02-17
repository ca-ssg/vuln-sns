FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y curl
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
