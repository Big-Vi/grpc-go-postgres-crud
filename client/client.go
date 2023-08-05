package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Big-Vi/realtime-dashboard-grpc-go-react/orderpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := orderpb.NewOrderServiceClient(conn)

	// Create Order
	fmt.Println("Creating order req")

	order := &orderpb.Order{
		CustomerId: "666",
		ProductName: "Test",
		Price: 999,
	}
	createOrderRes, err := c.CreateOrder(context.Background(), &orderpb.CreateOrderRequest{Order: order})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(createOrderRes)
}