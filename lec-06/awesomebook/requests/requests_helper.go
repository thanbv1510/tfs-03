package requests

import "awesomebook/entities"

func (bookRequest BookRequest) BookModel2BookEntity() entities.BookEntity {
	return entities.BookEntity{
		ID:       bookRequest.ID,
		Name:     bookRequest.Name,
		Author:   bookRequest.Author,
		Category: bookRequest.Category,
		Price:    bookRequest.Price,
	}
}
