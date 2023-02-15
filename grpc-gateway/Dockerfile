FROM golang:1.20-alpine AS build

RUN apk add --update --no-cache git
WORKDIR /app
# ENV GOPATH=/app

# Copy all of the staged files (protos plus go source)
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . /app/

# Download the go dependencies.

WORKDIR /app
# Build the gateway
RUN go build -o grpc_gateway

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/grpc_gateway /app/

EXPOSE 8081
CMD ["/app/grpc_gateway"]
