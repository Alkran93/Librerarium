package handlers

import (
	"cart-service/db"
	"cart-service/models"
	"encoding/json"
	"net/http"
)

func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT product_id, quantity FROM cart_items")
	if err != nil {
		http.Error(w, "Failed to fetch cart", 500)
		return
	}
	defer rows.Close()

	var cart []models.CartItem
	for rows.Next() {
		var item models.CartItem
		rows.Scan(&item.ProductID, &item.Quantity)
		cart = append(cart, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	var item models.CartItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil || item.ProductID == "" || item.Quantity <= 0 {
		http.Error(w, "Invalid item", 400)
		return
	}

	_, err = db.DB.Exec("INSERT INTO cart_items (product_id, quantity) VALUES (?, ?)", item.ProductID, item.Quantity)
	if err != nil {
		http.Error(w, "Failed to add item", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "item added"})
}
