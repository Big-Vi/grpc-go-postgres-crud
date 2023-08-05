package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Big-Vi/realtime-dashboard-grpc-go-react/db"
	"github.com/Big-Vi/realtime-dashboard-grpc-go-react/orderpb"
	"google.golang.org/grpc"
	"github.com/joho/godotenv"
)

type server struct {
	orderpb.OrderServiceServer
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
			Id: "5559",
			CustomerId: order.GetCustomerId(),
			ProductName: order.GetProductName(),
			Price: order.GetPrice(),
		},
	}, nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	_, err = db.InitDB()
	if err != nil {
		log.Fatalf("DB connection went wrong: %v", err)
		os.Exit(1)
	}

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