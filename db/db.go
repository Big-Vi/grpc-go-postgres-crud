package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Big-Vi/realtime-dashboard-grpc-go-react/types"
	"github.com/jackc/pgx/v5"
)

const (
	DbTimeout = 40
)

var DBClient = DBImpl{}

type DBImpl struct {
	conn *pgx.Conn
}

func InitDB() (DBImpl, error) {
	DATABASE_URL := getCoonectionString() 
	
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return DBImpl{}, err
	}

	DBClient = DBImpl{conn: conn}
	return DBImpl{conn: conn}, nil
}

func (db *DBImpl) CreateOrder(order *types.OrderItem) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout * time.Second)
	defer cancel()

	query := `INSERT INTO orders (customer_id, product_name, price, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int32
	err := DBClient.conn.QueryRow(ctx, query, order.CustomerID, order.ProductName, order.Price, order.Quantity, order.CreatedAt, order.UpdatedAt).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return id, nil
}

func (db *DBImpl) UpdateOrder(order_id int32, order *types.OrderItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout * time.Second)
	defer cancel()

	query := `UPDATE orders set quantity = $2, updated_at = $3 WHERE id = $1`

	_, err := DBClient.conn.Exec(ctx, query, order_id, order.Quantity, order.UpdatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Update order failed: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func (db *DBImpl) ReadOrder(order_id int32) (*types.OrderItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout * time.Second)
	defer cancel()

	query := `SELECT * FROM orders WHERE id = $1`

	order := types.OrderItem{}
	err := DBClient.conn.QueryRow(ctx, query, order_id).Scan(&order.ID, &order.CustomerID, &order.ProductName, &order.Price, &order.Quantity, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading order failed: %v\n", err)
		os.Exit(1)
	}
	return &order, nil
}

func (db *DBImpl) DeleteOrder(order_id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout * time.Second)
	defer cancel()

	query := `DELETE from orders WHERE id = $1`

	_, err := DBClient.conn.Exec(ctx, query, order_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading order failed: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func (db *DBImpl) ListAllOrders() ([]*types.OrderItem ,error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout * time.Second)
	defer cancel()

	query := `SELECT * FROM orders`

	orders := []*types.OrderItem{}
	rows, err := DBClient.conn.Query(ctx, query)
	for rows.Next() {
		order := types.OrderItem{}
		err := rows.Scan(&order.ID, &order.CustomerID, &order.ProductName, &order.Price, &order.Quantity, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Reading order failed: %v\n", err)
			os.Exit(1)
		}
		orders = append(orders, &order)
	}
	return orders, err
}

func getCoonectionString() string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	)
}
