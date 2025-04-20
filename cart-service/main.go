package main

import (
	"cart-service/db"
	"cart-service/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	router := mux.NewRouter()
	router.HandleFunc("/cart", handlers.GetCartHandler).Methods("GET")
	router.HandleFunc("/cart/add", handlers.AddToCartHandler).Methods("POST")
	router.HandleFunc("/cart/checkout", handlers.CheckoutHandler).Methods("POST")

	log.Println("Cart service with SQLite + JWT + MOM running on :" + os.Getenv("CART_PORT"))
	log.Fatal(http.ListenAndServe((":" + os.Getenv("CART_PORT")), router))
}
