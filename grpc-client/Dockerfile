FROM golang:1.20-alpine AS build

RUN apk add --update --no-cache git
WORKDIR /app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . /app/


WORKDIR /app
# Build the grpc_client
RUN go build -o grpc_client

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/grpc_client /app/

CMD ["/app/grpc_client"]
