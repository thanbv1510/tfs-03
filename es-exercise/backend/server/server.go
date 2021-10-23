package server

import (
	"awesomeProject/es"
	"awesomeProject/repository"
	"awesomeProject/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

const url = "http://localhost:9200"

func InitServer() error {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/search").Queries("keyword", "{keyword}").HandlerFunc(FindDataHandle)

	handler := cors.AllowAll().Handler(r)
	err := http.ListenAndServe(":8088", handler)
	if err != nil {
		return err
	}
	return nil
}

func FindDataHandle(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")

	esClient, _ := es.NewESClient(url)

	response := repository.FindData(keyword, esClient.Client)

	utils.ResponseWithJson(w, http.StatusOK, response)
}
