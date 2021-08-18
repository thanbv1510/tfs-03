package handlers

import (
	"awesomebook/entities"
	"awesomebook/helpers"
	"awesomebook/repositories"
	"awesomebook/requests"
	"encoding/json"
	"net/http"
)

func AddItemToOrderHandler(writer http.ResponseWriter, request *http.Request) {
	// Get userID from jwt
	// if not login then required login

	// Fix userID = 1
	userID := 1

	// Get shopping cartEntity from UserID
	cartEntity, err := repositories.GetCartFromUserID(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	itemRequest := &requests.ItemRequest{}
	err = json.NewDecoder(request.Body).Decode(itemRequest)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookExist, err := repositories.FindBookByID(itemRequest.BookID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartEntity.UserID = userID
	cartEntity.BookEntities = []entities.BookEntity{bookExist}
	cartEntity, err = repositories.SaveCart(cartEntity)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartBook := entities.CartBook{BookID: bookExist.ID,
		CartID: cartEntity.ID, Qty: itemRequest.Qty}

	helpers.DBConn().Save(cartBook)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = writer.Write([]byte("success!"))
}

func CartInfoHandler(writer http.ResponseWriter, request *http.Request) {
	// Get userID from jwt
	// if not login then required login

	// Fix userID = 1
	userID := 1

	// Get shopping cartEntity from UserID
	cartEntity, err := repositories.GetCartFromUserID(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.DBConn().Model(&cartEntity).Association("BookEntities").Find(&cartEntity.BookEntities)

	data, _ := repositories.FindAllCartBookByCartID(cartEntity.ID)

	resultJson, err := json.Marshal(cartEntity.CartEntity2CartInfoResponse(data))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(resultJson)
}

