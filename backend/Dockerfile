# Dockerfile.api
FROM golang:1.22.2-alpine3.18

WORKDIR /app

RUN apk --no-cache add bash git make gcc musl-dev curl
RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh
RUN make goose-up
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN go mod download

