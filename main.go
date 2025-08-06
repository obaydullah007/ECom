package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Obaydullah. I'm Software Engineer.")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Please request with GET method", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Please request with POST method", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to decode JSON", 400)
		return
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	w.WriteHeader(201)

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/create-products", createProduct)

	fmt.Println("Server running on port : 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Telescope",
		Description: "A Powerful Newtonian Telescope",
		Price:       199.99,
		ImgUrl:      "https://i.etsystatic.com/10189331/r/il/204377/3823588808/il_570xN.3823588808_cyr2.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Bike",
		Description: "A Cafe Racer Bike",
		Price:       2999.99,
		ImgUrl:      "https://caferacerdreams.es/wp-content/uploads/2016/02/crd11-cafe-racer-crd-honda-cb750kz-portada.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Laptop",
		Description: "A High End Gaming Laptop",
		Price:       1499.99,
		ImgUrl:      "https://gagadget.com/media/cache/18/f9/18f934fbff1cdb6f3ff5949b1c3ab9c5.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Smartphone",
		Description: "A Flagship Smartphone",
		Price:       999.99,
		ImgUrl:      "https://assets.gadgetandgear.com/upload/product/20240123_1705987892_224395.jpeg",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Jeep",
		Description: "A Powerful Off-Road Jeep",
		Price:       49999.99,
		ImgUrl:      "https://5.imimg.com/data5/SELLER/Default/2020/9/PW/NC/HD/90810037/willy-open-jeep.jpg",
	}
	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}
