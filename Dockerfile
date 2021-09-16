FROM golang:1.17.1

WORKDIR /app

COPY . .

RUN go mod download

#CMD ["go","run","./cmd/app/main.go"]
