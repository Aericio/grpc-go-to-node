package main

import (
	"context"
	pb "github.com/aericio/grpc/go-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("error")
		}
	}(conn)
	c := pb.NewServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	file, err := os.ReadFile("./payload.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	create, err := c.Create(ctx, &pb.CreateRequest{Key: "123", Value: file})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Request was created. %v", create.String())
}
