package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product Model
type Product struct {
	ID    string `json:"id"`
	Other *Other `json: other`
}

type Other struct {
	ID string `json:"id"`
}

// Mock Data

var products []Product

// Get All Products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typoe", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func main() {
	// Init Router
	r := mux.NewRouter()

	var other = Other{ID: "8675309"}

	// Mock Data - TODO: implement DB
	products = append(products, Product{ID: "1", Other: &other})
	products = append(products, Product{ID: "2", Other: &other})
	products = append(products, Product{ID: "3", Other: &other})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")

	// Start server, log if fatal
	log.Fatal(http.ListenAndServe(":8000", r))
}
