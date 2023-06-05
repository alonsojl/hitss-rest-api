FROM golang:1.17 AS build
WORKDIR /home/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify 
COPY . .
RUN openssl genrsa -out app.rsa 1024 && \
    openssl rsa -in app.rsa -pubout >  app.rsa.pub
RUN go build -o cmd/hitss cmd/main.go
CMD ["./cmd/hitss"]

