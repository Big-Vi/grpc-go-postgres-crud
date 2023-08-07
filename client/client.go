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

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
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

	createOrder(c)
	updateOrder(c)
	readOrder(c)
	listAllOrders(c)
	deleteOrder(c)
}

func createOrder(c orderpb.OrderServiceClient) {
	fmt.Println("Req for creating order")

	order := &orderpb.Order{
		CustomerId: "666",
		ProductName: "Testing",
		Price: 999,
		Quantity: 1,
	}
	createOrderRes, err := c.CreateOrder(context.Background(), &orderpb.CreateOrderRequest{Order: order})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(createOrderRes)
}

func updateOrder(c orderpb.OrderServiceClient) {
	fmt.Println("Req for updating order")
	order := &orderpb.Order{
		Id: 1,
		Quantity: 5,
	}
	updateOrderRes, err := c.UpdateOrder(context.Background(), &orderpb.UpdateOrderRequest{Order: order})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(updateOrderRes)
}

func readOrder(c orderpb.OrderServiceClient) {
	fmt.Println("Req for reading order")
	readOrderRes, err := c.ReadOrder(context.Background(), &orderpb.ReadOrderRequest{OrderId: 1})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(readOrderRes)
}

func listAllOrders(c orderpb.OrderServiceClient) {
	fmt.Println("Req for listing all orders")
	listAllOrdersRes, err := c.ListAllOrders(context.Background(), &orderpb.ListAllOrdersRequest{})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(listAllOrdersRes)
}

func deleteOrder(c orderpb.OrderServiceClient) {
	fmt.Println("Req for deleting order")
	deleteOrderRes, err := c.DeleteOrder(context.Background(), &orderpb.DeleteOrderRequest{OrderId: 1})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(deleteOrderRes)
}
