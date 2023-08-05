package types

type OrderItem struct {
	ID string `json:"id"`
	CustomerID string `json:"customer_id"`
	ProductName string `json:"product_name"`
	Price int32 `json:"price"`
}