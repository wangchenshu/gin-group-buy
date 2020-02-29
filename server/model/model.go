package model

// Product -
type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Link   string `json:"link"`
	PicUrl string `json:"pic_url"`
	Active string `json:"active"`
}

// Cart -
type Cart struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	LineUserID  string `json:"line_user_id"`
	Username    string `json:"username"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
}

// Order -
type Order struct {
	ID         int    `json:"id"`
	ProductID  int    `json:"product_id"`
	LineUserID string `json:"line_user_id"`
	Qty        int    `json:"qty"`
}
