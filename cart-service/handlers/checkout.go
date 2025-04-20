package handlers

import (
	"cart-service/db"
	"cart-service/models"
	"cart-service/mom"
	"encoding/json"
	"log"
	"net/http"
)

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT product_id, quantity FROM cart_items")
	if err != nil {
		http.Error(w, "Failed to read cart", 500)
		return
	}
	defer rows.Close()

	var cart []models.CartItem
	for rows.Next() {
		var item models.CartItem
		rows.Scan(&item.ProductID, &item.Quantity)
		cart = append(cart, item)
	}

	err = mom.PublishCheckout(cart)
	if err != nil {
		http.Error(w, "Failed to send MOM event", 500)
		return
	}

	_, err = db.DB.Exec("DELETE FROM cart_items")
	if err != nil {
		http.Error(w, "Failed to clear cart", 500)
		return
	}

	log.Println("Checkout completed: cart cleared and event sent")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "order placed"})
}
