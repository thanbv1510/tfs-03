package handlers

import (
	"awesomebook/repositories"
	"awesomebook/requests"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SaveBookHandler(writer http.ResponseWriter, request *http.Request) {
	bookRequest := &requests.BookRequest{}

	err := json.NewDecoder(request.Body).Decode(bookRequest)
	if err != nil {
		sugar.Error("Cannot decode book")
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	book, err := repositories.SaveBook(bookRequest.BookModel2BookEntity())
	if err != nil {
		sugar.Error("Cannot create book")
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("Cannot create book!"))

		return
	}

	fmt.Println(book)
	_, _ = writer.Write([]byte("success!"))
}

func GetAllBookHandler(writer http.ResponseWriter, request *http.Request) {
	result, err := repositories.FindAllBook()

	if err != nil {
		sugar.Error("Cannot get all book")
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("Cannot get all book!"))

		return
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(resultJson)

}

func GetBookHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, ok := params["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	idInt, _ := strconv.Atoi(id)

	book, err := repositories.FindBookByID(idInt)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	resultJson, err := json.Marshal(book)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(resultJson)
}

func DeleteBookHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, ok := params["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	idInt, _ := strconv.Atoi(id)
	book, err := repositories.FindBookByID(idInt)
	if err != nil || book.ID == 0 {
		sugar.Errorf("Book with ID = %d not exist!", idInt)
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf("Book with ID = %d not exist!", idInt)))

		return
	}

	err = repositories.DeleteBookByID(book.ID)
	if err != nil {
		sugar.Errorf("Book with ID = %d delete fail!, err = %s", idInt, err)
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf("Book with ID = %d delete fail!", idInt)))

		return
	}

	_, _ = writer.Write([]byte("Delete book success!"))
}
