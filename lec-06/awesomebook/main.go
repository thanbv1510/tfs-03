package main

import (
	"awesomebook/entities"
	"awesomebook/handlers"
	"awesomebook/helpers"
	"awesomebook/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var serveMux *mux.Router

func main() {
	// Public API
	publicRoute := serveMux.Methods(http.MethodPost).Subrouter()
	publicRoute.HandleFunc("/signup", handlers.RegisterHandler)
	publicRoute.HandleFunc("/login", handlers.LoginHandler)

	// Middleware Public API
	publicRoute.Use(middleware.ContentTypeCheckingMiddleware)

	// Private API
	// POST and PUT
	privateRouteP2 := serveMux.NewRoute().Subrouter()
	// Product management (Book)
	privateRouteP2.HandleFunc("/books", handlers.SaveBookHandler).Methods(http.MethodPost)
	privateRouteP2.HandleFunc("/books", handlers.SaveBookHandler).Methods(http.MethodPut)

	// Shopping cart management
	privateRouteP2.HandleFunc("/cart/items", handlers.AddItemToOrderHandler).Methods(http.MethodPost)

	// Payment
	privateRouteP2.HandleFunc("/payment", handlers.PaymentHandler).Methods(http.MethodPost)

	// Middleware POST and PUT
	privateRouteP2.Use(middleware.ContentTypeCheckingMiddleware)

	// GET and DELETE
	privateRouteGD := serveMux.NewRoute().Subrouter()
	// Product management (Book)
	privateRouteGD.HandleFunc("/books", handlers.GetAllBookHandler).Methods(http.MethodGet)
	privateRouteGD.HandleFunc("/books/{id:[0-9]+}", handlers.GetBookHandler).Methods(http.MethodGet)
	privateRouteGD.HandleFunc("/books/{id:[0-9]+}", handlers.DeleteBookHandler).Methods(http.MethodDelete)

	privateRouteGD.HandleFunc("/carts", handlers.CartInfoHandler).Methods(http.MethodGet)

	// Middleware GET and DELETE

	// Server
	server := &http.Server{
		Handler:      serveMux,
		Addr:         "localhost:8901",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
		return
	}
}

func init() {
	err := helpers.DBConn().AutoMigrate(&entities.UserEntity{}, &entities.BookEntity{}, &entities.OrderEntity{}, &entities.PaymentEntity{}, &entities.CartEntity{}, &entities.CartBook{}, &entities.OrderBook{})
	if err != nil {
		panic(err)
	}

	err = helpers.DBConn().SetupJoinTable(&entities.CartEntity{}, "BookEntities", &entities.CartBook{})
	if err != nil {
		panic(err)
	}

	err = helpers.DBConn().SetupJoinTable(&entities.OrderEntity{}, "BookEntities", &entities.OrderBook{})
	if err != nil {
		panic(err)
	}

	// Create a serve mux
	serveMux = mux.NewRouter()
}
