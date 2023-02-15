package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-client/pkg"
	"log"
	"sync"
	"time"
)

func main() {
	target := "bench-server:6790"
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("error request", err)
	}
	defer conn.Close()

	start := time.Now()
	server := pkg.NewCoreServiceClient(conn)
	lock := &sync.WaitGroup{}

	total := int64(50000)
	for i := int64(0); i < total; i++ {
		lock.Add(1)
		go func() {
			request := &pkg.PingRequest{}
			_, err := server.Ping(context.Background(), request)
			if err != nil {
				log.Print(err)
			}
			lock.Done()
		}()
	}
	lock.Wait()

	execTime := time.Now().Sub(start)
	fmt.Printf("Benchmark result for %d request in %0.2f ms\n", total, execTime.Seconds())
	fmt.Printf("Throughtput : %d req/s\n", total*1000000/execTime.Microseconds())
	fmt.Printf("Latency : %d ms\n", execTime.Microseconds()/total)
}
