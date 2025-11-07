package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Fprintf(w, "Halo Aku Golang!")
}

func produkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Fprintf(w, "Selamat datang di halam produk!")
}

func main() {
	http.HandleFunc("/halo", helloHandler)
	http.HandleFunc("/produk", produkHandler)
	fmt.Println("Auth Service Running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
