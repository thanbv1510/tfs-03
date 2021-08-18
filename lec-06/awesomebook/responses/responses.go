package responses

type CartInfoResponse struct {
	CartID int            `json:"cart_id"`
	UserID int            `json:"user_id"`
	Books  []BookResponse `json:"books"`
}

type BookResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Qty      int    `json:"qty"`
}
