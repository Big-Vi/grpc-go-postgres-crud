package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Big-Vi/realtime-dashboard-grpc-go-react/orderpb"
	"google.golang.org/grpc"
)

type server struct {
	orderpb.OrderServiceServer
}

type OrderItem struct {
	ID string `json:"id"`
	CustomerID string `json:"customer_id"`
	ProductName string `json:"product_name"`
	Price int32 `json:"price"`
}

func (s *server) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error){
	fmt.Println("Creating order....")
	order := req.GetOrder()

	// orderData := OrderItem {
	// 	ID: "555",
	// 	CustomerID: order.GetCustomerId(),
	// 	ProductName: order.GetProductName(),
	// 	Price: order.GetPrice(),
	// }
	return &orderpb.CreateOrderResponse{
		Order: &orderpb.Order{
			Id: "555",
			CustomerId: order.GetCustomerId(),
			ProductName: order.GetProductName(),
			Price: order.GetPrice(),
		},
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("test...")
	s := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}