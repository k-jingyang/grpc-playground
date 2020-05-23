package main

import (
	"context"
	pb "grpc-playground/playground"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	const port = "localhost:50051"

	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("Unable to dial %v", port)
	}
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Hello from client"})
	if err != nil {
		log.Fatalf("Unable to say hello: %v", err)
	}
	log.Printf("From server: %v", reply.GetMessage())
}
