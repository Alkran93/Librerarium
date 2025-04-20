package main

import (
	"database/sql"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
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

var jwtSecret = []byte("super-secret-key")

func getUsernameFromToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", http.ErrNoCookie // usar cualquier error
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	username, ok := claims["sub"].(string)
	if !ok {
		return "", err
	}

	return username, nil
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	username, err := getUsernameFromToken(r)
	if err != nil {
		http.Error(w, "Unauthorized or invalid token", http.StatusUnauthorized)
		return
	}

	rows, err := db.Query("SELECT product_id, quantity FROM cart_items")
	if err != nil {
		http.Error(w, "Error leyendo carrito", 500)
		return
	}
	defer rows.Close()

	var cart []CartItem
	for rows.Next() {
		var item CartItem
		rows.Scan(&item.ProductID, &item.Quantity)
		cart = append(cart, item)
	}

	// Enviar a RabbitMQ
	conn, err := amqp.Dial("amqp://user:password@3.82.109.178:5672/")
	if err != nil {
		log.Fatal("[Error] No se pudo conectar a RabbitMQ:", err)
	}
	defer conn.Close()
	ch, _ := conn.Channel()
	defer ch.Close()

	ch.ExchangeDeclare("my_exchange", "direct", true, false, false, false, nil)

	body, _ := json.Marshal(map[string]interface{}{
		"evento":  "checkout",
		"usuario": username,
		"items":   cart,
	})

	ch.Publish("my_exchange", "test", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	// Limpiar carrito
	_, err = db.Exec("DELETE FROM cart_items")
	if err != nil {
		http.Error(w, "Failed to checkout", 500)
		return
	}

	log.Printf("Checkout de %s enviado a MOM y carrito limpiado.\n", username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "order placed"})
}
