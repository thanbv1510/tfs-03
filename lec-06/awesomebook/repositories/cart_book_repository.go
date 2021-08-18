package repositories

import (
	"awesomebook/entities"
	"awesomebook/helpers"
)

func FindAllCartBookByCartID(cartID int) ([]entities.CartBook, error) {
	var cartBooks []entities.CartBook
	err := helpers.DBConn().Find(&cartBooks).Where("cart_entity_id = ?", cartID).Error
	return cartBooks, err
}
