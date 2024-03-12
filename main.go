package main

import (
	"context"
	"log"
)

func main() {
	addr := "localhost:5436"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	req := &pb.HelloRequest{
		Message: "sosi",
	}

	resp, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %s", resp.Message)
}
