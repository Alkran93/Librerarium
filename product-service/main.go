package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       float64 `json:"price"`
	Available   int    `json:"available"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite", "./products.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT,
			price REAL,
			available INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products", addProduct).Methods("POST")

	log.Println("Product service running on port 3003")
	log.Fatal(http.ListenAndServe(":3003", router))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, author, price, available FROM products")
	if err != nil {
		http.Error(w, "DB query failed", 500)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Title, &p.Author, &p.Price, &p.Available)
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil || p.Title == "" || p.Price <= 0 {
		http.Error(w, "Invalid product data", 400)
		return
	}

	_, err = db.Exec("INSERT INTO products (title, author, price, available) VALUES (?, ?, ?, ?)",
		p.Title, p.Author, p.Price, p.Available)

	if err != nil {
		http.Error(w, "Insert failed", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "product added"})
}
