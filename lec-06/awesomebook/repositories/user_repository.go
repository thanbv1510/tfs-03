package repositories

import (
	"awesomebook/entities"
	"awesomebook/helpers"
)

func CreateUser(user entities.UserEntity) {
	helpers.DBConn().Create(&user)
}

func FindByUsername(username string) (user entities.UserEntity) {
	helpers.DBConn().Where("username = ?", username).First(&user)
	return
}

func FindUserByUserID(userID int) (user entities.UserEntity, err error) {
	err = helpers.DBConn().Where("id = ?", userID).First(&user).Error
	return
}
