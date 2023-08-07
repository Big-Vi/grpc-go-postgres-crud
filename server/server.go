package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/Big-Vi/realtime-dashboard-grpc-go-react/db"
	"github.com/Big-Vi/realtime-dashboard-grpc-go-react/orderpb"
	"github.com/Big-Vi/realtime-dashboard-grpc-go-react/types"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	orderpb.OrderServiceServer
	db db.DBImpl
}

var database db.DBImpl

func (s *server) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error){
	fmt.Println("Creating order....")
	order := req.GetOrder()

	orderData := types.OrderItem {
		CustomerID: order.GetCustomerId(),
		ProductName: order.GetProductName(),
		Price: order.GetPrice(),
		Quantity: order.GetQuantity(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	id, err := database.CreateOrder(&orderData)
	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{
		Order: &orderpb.Order{
			Id: id,
			CustomerId: order.GetCustomerId(),
			ProductName: order.GetProductName(),
			Price: order.GetPrice(),
			Quantity: order.GetQuantity(),
		},
	}, nil
}

func (s *server) UpdateOrder(ctx context.Context, req *orderpb.UpdateOrderRequest) (*orderpb.UpdateOrderResponse, error) {
	fmt.Println("Updating order...")
	order := req.GetOrder()
	order_id := order.GetId()
	orderData := types.OrderItem {
		Quantity: order.GetQuantity(),
		UpdatedAt: time.Now().UTC(),
	}

	err := database.UpdateOrder(order_id, &orderData)
	if err != nil {
		return nil, err
	}

	return &orderpb.UpdateOrderResponse{
		Order: &orderpb.Order{
			Id: order_id,
			CustomerId: order.GetCustomerId(),
			ProductName: order.GetProductName(),
			Price: order.GetPrice(),
			Quantity: order.GetQuantity(),
		},
	}, nil
}

func(s *server) ReadOrder(ctx context.Context, req *orderpb.ReadOrderRequest) (*orderpb.ReadOrderResponse, error) {
	fmt.Println("Reading order...")
	order_id := req.GetOrderId()

	orderData, err := database.ReadOrder(order_id)
	fmt.Println(orderData)
	if err != nil {
		return nil, err
	}

	return &orderpb.ReadOrderResponse{
		Order: &orderpb.Order{
			Id: order_id,
			CustomerId: orderData.CustomerID,
			ProductName: orderData.ProductName,
			Price: orderData.Price,
			Quantity: orderData.Quantity,
		},
	}, nil
}

func (s *server) DeleteOrder(ctx context.Context, req *orderpb.DeleteOrderRequest) (*orderpb.DeleteOrderResponse, error) {
	fmt.Println("Deleting order...")
	order_id := req.GetOrderId()

	err := database.DeleteOrder(order_id)
	if err != nil {
		return nil, err
	}

	return &orderpb.DeleteOrderResponse{
		OrderId: order_id,
	}, nil
}

func (s *server) ListAllOrders(ctx context.Context, req *orderpb.ListAllOrdersRequest) (*orderpb.ListAllOrdersResponse, error) {
	fmt.Println("Listing all orders")

	ordersData, err := database.ListAllOrders()
	if err != nil {
		return nil, err
	}

	var orders []*orderpb.Order

	for _, ele := range ordersData {
		order := orderpb.Order{
			Id: int32(ele.ID),
			CustomerId: ele.CustomerID,
			ProductName: ele.ProductName,
			Price: ele.Price,
			Quantity: ele.Quantity,
		}
		orders = append(orders, &order)
	}

	return &orderpb.ListAllOrdersResponse{
		Orders: orders,
	}, nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	database, err = db.InitDB()
	if err != nil {
		log.Fatalf("DB connection went wrong: %v", err)
		os.Exit(1)
	}

	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := &server{}
	orderpb.RegisterOrderServiceServer(grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}