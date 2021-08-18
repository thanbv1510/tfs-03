package requests

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Username string `json:"username"`
	Passwd   string `json:"password"`
	Email    string `json:"email"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type BookRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

type ItemRequest struct {
	BookID int `json:"book_id"`
	Qty    int `json:"qty"`
}

type ChargeRequest struct {
	Amount       int64  `json:"amount"`
	Username     string `json:"username"`
	ReceiptEmail string `json:"receiptEmail"`
}
