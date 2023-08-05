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

type DBImpl struct {
	conn *pgx.Conn
}

func InitDB() (DBImpl, error) {
	DATABASE_URL := getCoonectionString() 
	fmt.Println(DATABASE_URL)
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return DBImpl{}, err
	}
	defer conn.Close(context.Background())

	return DBImpl{conn: conn}, nil
}

func (db *DBImpl) CreateOrder(order *types.OrderItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout * time.Second)
	defer cancel()

	query := `INSERT INTO orders (customer_id, product_name, price) VALUES ($!, $2, $3) RETURNING id`
	stmt, err := db.conn.Prepare(ctx, "insert_order", query)
	if err != nil {
		return err
	}

	rows, err := db.conn.Query(ctx, stmt.Name, order.CustomerID, order.ProductName, order.Price)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&order.ID)
		if err != nil {
			return err
		}
	}

	return nil
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
