FROM golang:1.20-alpine AS build

RUN apk add --update --no-cache git
WORKDIR /app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . /app/


WORKDIR /app
# Build the grpc-server
RUN go build -o grpc-server

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/grpc-server /app/

CMD ["/app/grpc-server"]
