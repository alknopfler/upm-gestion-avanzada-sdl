# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM docker.io/library/golang:1.19.2-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN apk add build-base

# build an app
COPY . .
RUN go build -v -o /mybankapi

CMD [ "/mybankapi" ]