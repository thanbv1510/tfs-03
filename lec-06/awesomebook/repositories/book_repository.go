package repositories

import (
	"awesomebook/entities"
	"awesomebook/helpers"
)

func SaveBook(book entities.BookEntity) (entities.BookEntity, error) {
	err := helpers.DBConn().Save(&book).Error
	return book, err
}

func FindAllBook() (result []entities.BookEntity, err error) {
	err = helpers.DBConn().Find(&result).Error
	return
}

func FindBookByID(id int) (book entities.BookEntity, err error) {
	err = helpers.DBConn().First(&book, id).Error
	return
}

func DeleteBookByID(id int) error {
	return helpers.DBConn().Delete(&entities.BookEntity{}, id).Error
}
