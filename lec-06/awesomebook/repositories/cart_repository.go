package repositories

import (
	"awesomebook/entities"
	"awesomebook/helpers"
)

func GetCartFromUserID(userID int) (cart entities.CartEntity, err error) {
	var result entities.CartEntity
	db := helpers.DBConn()
	err = db.First(&cart, userID).Error
	if err != nil {
		result.UserID = userID
		err = db.Save(&result).Error
		return result, err
	}

	return
}

func SaveCart(cart entities.CartEntity) (entities.CartEntity, error) {
	err := helpers.DBConn().Save(&cart).Error
	return cart, err
}

func GetAllCart(cartEntity entities.CartEntity) entities.CartEntity {
	helpers.DBConn().Model(&cartEntity).Association("BookEntities").Find(&cartEntity.BookEntities)
	return cartEntity
}
