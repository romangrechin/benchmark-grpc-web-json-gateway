package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"

	"grpc-server/pkg"
)

var (
	grpcPort = flag.Int("grpc-port", 6790, "gRPC port")
	httpPort = flag.Int("http-port", 6789, "HTTP port")
)

// server is used to implement pkg.CoreServiceServer.
type server struct {
	pkg.UnimplementedCoreServiceServer
}

// Ping implements pkg.CoreServiceServer
func (s *server) Ping(ctx context.Context, in *pkg.PingRequest) (*pkg.PingResponse, error) {
	return &pkg.PingResponse{Message: "pong", Timestamp: time.Now().Unix()}, nil
}

func serveGRPC(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}
	s := grpc.NewServer()
	pkg.RegisterCoreServiceServer(s, &server{})
	log.Printf("gRPC server listening at :%d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to gRPC serve: %v", err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		d := json.NewDecoder(r.Body)
		p := &pkg.PingRequest{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		resp := &pkg.PingResponse{Message: "pong", Timestamp: time.Now().Unix()}
		json.NewEncoder(w).Encode(resp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func serveHTTP(port int) {
	http.HandleFunc("/ping", httpHandler)
	log.Printf("HTTP server listening at :%d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("failed to HTTP serve: %v", err)
	}
}

func main() {
	flag.Parse()
	go serveGRPC(*grpcPort)
	serveHTTP(*httpPort)
}
