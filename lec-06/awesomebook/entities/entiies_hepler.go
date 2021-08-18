package entities

import "awesomebook/responses"

func (cartEntity CartEntity) CartEntity2CartInfoResponse(cartBooks []CartBook) responses.CartInfoResponse {
	var result responses.CartInfoResponse
	dataMap := make(map[int]int)
	for _, v := range cartBooks {
		dataMap[v.BookID] = v.Qty
	}

	result.CartID = cartEntity.ID
	result.UserID = cartEntity.UserID

	var books []responses.BookResponse
	for _, book := range cartEntity.BookEntities {
		qty, _ := dataMap[book.ID]
		books = append(books, responses.BookResponse{
			ID:       book.ID,
			Name:     book.Name,
			Category: book.Category,
			Author:   book.Author,
			Price:    book.Price,
			Qty:      qty,
		})

	}
	result.Books = books
	return result
}
