package handlers

import (
	"cart-service/db"
	"cart-service/models"
	"cart-service/mom"
	"cart-service/util"
	"encoding/json"
	"log"
	"net/http"
)

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	username, err := util.GetUsernameFromToken(r)
	if err != nil {
		http.Error(w, "Unauthorized or invalid token", http.StatusUnauthorized)
		return
	}

	rows, err := db.DB.Query("SELECT product_id, quantity FROM cart_items")
	if err != nil {
		http.Error(w, "Error reading cart", 500)
		return
	}
	defer rows.Close()

	var cart []models.CartItem
	for rows.Next() {
		var item models.CartItem
		rows.Scan(&item.ProductID, &item.Quantity)
		cart = append(cart, item)
	}

	err = mom.PublishCheckout(username, cart)
	if err != nil {
		http.Error(w, "Failed to send checkout to MOM", 500)
		return
	}

	_, err = db.DB.Exec("DELETE FROM cart_items")
	if err != nil {
		http.Error(w, "Failed to clear cart", 500)
		return
	}

	log.Printf("Checkout de %s enviado a MOM y carrito limpiado.\n", username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "order placed"})
}
