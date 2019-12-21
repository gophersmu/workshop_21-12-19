package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Product is a model for products
type Product struct {
	ID       int    `json:"id"`
	Product  string `json:"product"`
	Image    string `json:"image"`
	Quantity int    `json:"quantity"`
}

var products []Product

func init() {
	var id int

	id = len(products) + 1
	products = append(products, Product{
		ID:       id,
		Product:  "Product 1",
		Image:    "/images/products/1.jpg",
		Quantity: 1,
	})
	id = len(products) + 1
	products = append(products, Product{
		ID:       id,
		Product:  "Product 2",
		Image:    "/images/products/2.jpg",
		Quantity: 1,
	})
	id = len(products) + 1
	products = append(products, Product{
		ID:       id,
		Product:  "Product 3",
		Image:    "/images/products/3.jpg",
		Quantity: 1,
	})
	id = len(products) + 1
	products = append(products, Product{
		ID:       id,
		Product:  "Product 4",
		Image:    "/images/products/4.jpg",
		Quantity: 1,
	})
	id = len(products) + 1
	products = append(products, Product{
		ID:       id,
		Product:  "Product 5",
		Image:    "/images/products/5.jpg",
		Quantity: 1,
	})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/products", listProducts).Methods("GET")
	router.HandleFunc("/product", createProduct).Methods("POST")
	router.HandleFunc("/product/{id}", getProduct).Methods("GET")
	router.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")

	fmt.Println("Listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product

	json.NewDecoder(r.Body).Decode(&p)

	p.ID = len(products) + 1
	products = append(products, p)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id := getID(r)
	for _, p := range products {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var p Product

	id := getID(r)
	json.NewDecoder(r.Body).Decode(&p)

	for i := range products {
		if products[i].ID == id {
			products[i].Product = p.Product
			products[i].Image = p.Image
			products[i].Quantity = p.Quantity

			json.NewEncoder(w).Encode(products[i])
			return
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	var prods []Product

	id := getID(r)
	for _, p := range products {
		if p.ID != id {
			prods = append(prods, p)
		}
	}

	products = prods
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getID(r *http.Request) int {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	return i
}
