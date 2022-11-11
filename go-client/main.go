package main

import (
	"context"
	"fmt"
	pb "github.com/aericio/grpc/go-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:4362", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("error")
		}
	}(conn)
	c := pb.NewQueueClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	file, err := os.ReadFile("./payload.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	for i := 0; i < 5; i++ {
		start := time.Now()
		_, err := c.Push(ctx, &pb.EventRequest{Key: "123", Value: file})
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Printf("Request was created")
		duration := time.Since(start)
		fmt.Println(duration)
	}
}
