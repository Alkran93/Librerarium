package main

import (
	"cart-service/db"
	"cart-service/handlers"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()

	router := mux.NewRouter()
	router.HandleFunc("/cart", handlers.GetCartHandler).Methods("GET")
	router.HandleFunc("/cart/add", handlers.AddToCartHandler).Methods("POST")
	router.HandleFunc("/cart/checkout", handlers.CheckoutHandler).Methods("POST")

	log.Println("Cart service running on port 3002")
	log.Fatal(http.ListenAndServe(":3002", router))
}
