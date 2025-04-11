package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

var cart []CartItem

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/cart", getCartHandler).Methods("GET")
	router.HandleFunc("/cart/add", addToCartHandler).Methods("POST")
	router.HandleFunc("/cart/checkout", checkoutHandler).Methods("POST")

	log.Println("Cart service running on :3002")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func getCartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	var item CartItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Append to in-memory cart
	cart = append(cart, item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "item added"})
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate checkout
	log.Println("Checkout processed:", cart)
	cart = []CartItem{} // clear cart

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "order placed"})
}
