package handlers

import (
	"awesomebook/entities"
	"awesomebook/helpers"
	"awesomebook/repositories"
	"awesomebook/requests"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var sugar = helpers.GetSugar()

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	credentials := &requests.Credentials{}

	err := json.NewDecoder(request.Body).Decode(credentials)
	if err != nil {
		sugar.Error("Cannot decode user info")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userExist := repositories.FindByUsername(credentials.Username)
	if userExist.Username != "" {
		sugar.Errorf("Exist user with username: %s", userExist.Username)
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(fmt.Sprintf("Username: %s exist!", userExist.Username)))
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(credentials.Passwd), 8)
	if err != nil {
		sugar.Error("Cannot hash password")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userEntity := entities.UserEntity{Username: credentials.Username, Passwd: string(hashedPasswd), Email: credentials.Email, CartEntity: entities.CartEntity{}}
	repositories.CreateUser(userEntity)

	_, _ = writer.Write([]byte("Create user success!"))
}

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	credentials := &requests.Credentials{}

	err := json.NewDecoder(request.Body).Decode(credentials)
	if err != nil {
		sugar.Error("Cannot decode user info")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userExist := repositories.FindByUsername(credentials.Username)
	if userExist.Username == "" {
		sugar.Errorf("Not exist user with username: %s", userExist.Username)
		writer.WriteHeader(http.StatusUnauthorized)
		_, _ = writer.Write([]byte(fmt.Sprintf("Username: %s not exist!", userExist.Username)))
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(userExist.Passwd), []byte(credentials.Passwd)); err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		_, _ = writer.Write([]byte(fmt.Sprintf("Username: %s not exist!", userExist.Username)))
		return
	}

	// Generate token and set for response user
}
