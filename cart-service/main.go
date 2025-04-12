package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite", "./cart.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cart_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_id TEXT NOT NULL,
			quantity INTEGER NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/cart", getCartHandler).Methods("GET")
	router.HandleFunc("/cart/add", addToCartHandler).Methods("POST")
	router.HandleFunc("/cart/checkout", checkoutHandler).Methods("POST")

	log.Println("Cart service with SQLite running on :3002")
	log.Fatal(http.ListenAndServe(":3002", router))
}

func getCartHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT product_id, quantity FROM cart_items")
	if err != nil {
		http.Error(w, "Failed to fetch cart", 500)
		return
	}
	defer rows.Close()

	var cart []CartItem
	for rows.Next() {
		var item CartItem
		rows.Scan(&item.ProductID, &item.Quantity)
		cart = append(cart, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	var item CartItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil || item.ProductID == "" || item.Quantity <= 0 {
		http.Error(w, "Invalid item", 400)
		return
	}

	_, err = db.Exec("INSERT INTO cart_items (product_id, quantity) VALUES (?, ?)", item.ProductID, item.Quantity)
	if err != nil {
		http.Error(w, "Failed to add item", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "item added"})
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("DELETE FROM cart_items")
	if err != nil {
		http.Error(w, "Failed to checkout", 500)
		return
	}

	log.Println("Order placed and cart cleared.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "order placed"})
}
