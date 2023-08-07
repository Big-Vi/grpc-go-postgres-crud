package types

import "time"

type OrderItem struct {
	ID int `json:"id"`
	CustomerID string `json:"customer_id"`
	ProductName string `json:"product_name"`
	Price int32 `json:"price"`
	Quantity int32 `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}