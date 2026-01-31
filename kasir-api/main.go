package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"strconv"
	"kasir-api/database"
	"kasir-api/handlers"
	"os"
	"log"
	"kasir-api/repositories"
	"kasir-api/services"
	"github.com/spf13/viper"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

var product =[]Product{
	{ID:1, Name:"Indomie Goreng", Price:2000, Stock:10},
	{ID:2, Name:"Vit 1000ml", Price:9000, Stock:20},
	{ID:3, Name:"Susu Ultra", Price:10000, Stock:10},
	{ID:4, Name:"Sosis", Price:90000, Stock:7},
}

var category=[]Category{
	{ID:1, Name:"Mie", Description:"Aneka Mie"},
	{ID:2, Name:"Bumbu Dapur", Description:"Aneka Bumbu Dapur"},
	{ID:3, Name:"Obat", Description:"Obat-obatan untuk sakit apapun, termasuk sakit hati."},
} 

type Config struct {
	Port    string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func getProductByID(w http.ResponseWriter, r *http.Request){
		idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
			return
		}		

		for _, p := range product {
			if p.ID == id {
				w.Header().Set("Content-Type","application/json")
				json.NewEncoder(w).Encode(p)
				return
			}
		}

		http.Error(w, "Product Belum Ada", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i := range product {
		if product[i].ID == id {
			product[i] = updateProduct

			json.NewEncoder(w).Encode(updateProduct)
			return
		}
 	}
		
	http.Error(w, "Product Belum Ada", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request){
		idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
			return
		}

		for i, p := range product {
			if p.ID == id {
				product=append(product[:i], product[i+1:]...)
				w.Header().Set("Content-Type","application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"message":"Success Delete",
				})

				return
			}
		}

		http.Error(w, "Product Belum Ada", http.StatusNotFound)
}

func getCategoryByID(w http.ResponseWriter, r *http.Request){
		idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}		

		for _, c := range category {
			if c.ID == id {
				w.Header().Set("Content-Type","application/json")
				json.NewEncoder(w).Encode(c)
				return
			}
		}

		http.Error(w, "Category Belum Ada", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i := range category {
		if category[i].ID == id {
			category[i] = updateCategory

			json.NewEncoder(w).Encode(updateCategory)
			return
		}
 	}
		
	http.Error(w, "Category Belum Ada", http.StatusNotFound)
}

func deleteCategory(w http.ResponseWriter, r *http.Request){
		idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}

		for i, c := range category {
			if c.ID == id {
				category=append(category[:i], category[i+1:]...)
				w.Header().Set("Content-Type","application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"message":"Success Delete",
				})

				return
			}
		}

		http.Error(w, "Category Belum Ada", http.StatusNotFound)
}


func main(){
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	fmt.Println("Config:",config.DBConn)

	
	
	// http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request){
	// 	if r.Method == "GET" {
	// 		getProductByID(w, r)
	// 	}else if r.Method == "PUT" {
	// 		updateProduct(w, r)
	// 	}else if r.Method == "DELETE" {
	// 		deleteProduct(w, r)
	// 	}
		
	// })

	// http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request){
	// 	if r.Method == "GET" {
	// 		w.Header().Set("Content-Type","application/json")
	// 		json.NewEncoder(w).Encode(product)
	// 	}else if r.Method == "POST"{
	// 		var newProduct Product
	// 		err := json.NewDecoder(r.Body).Decode(&newProduct)
	// 		if err != nil {
	// 			http.Error(w, "Invalid request", http.StatusBadRequest)
	// 		}

	// 		newProduct.ID=len(product) + 1
	// 		product=append(product, newProduct)

	// 		w.Header().Set("Content-Type","application/json")
	// 		w.WriteHeader(http.StatusCreated)
	// 		json.NewEncoder(w).Encode(newProduct)
	// 	}
	// })

	// http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request){
	// 	if r.Method == "GET" {
	// 		getCategoryByID(w, r)
	// 	}else if r.Method == "PUT" {
	// 		updateCategory(w, r)
	// 	}else if r.Method == "DELETE" {
	// 		deleteCategory(w, r)
	// 	}
		
	// })

	// http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request){
	// 	if r.Method == "GET" {
	// 		w.Header().Set("Content-Type","application/json")
	// 		json.NewEncoder(w).Encode(category)
	// 	}else if r.Method == "POST"{
	// 		var newCategory Category
	// 		err := json.NewDecoder(r.Body).Decode(&newCategory)
	// 		if err != nil {
	// 			http.Error(w, "Invalid request", http.StatusBadRequest)
	// 		}

	// 		newCategory.ID=len(category) + 1
	// 		category=append(category, newCategory)

	// 		w.Header().Set("Content-Type","application/json")
	// 		w.WriteHeader(http.StatusCreated)
	// 		json.NewEncoder(w).Encode(newCategory)
	// 	}
	// })

	// http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
	// 	w.Header().Set("Content-Type","application/json")
	// 	json.NewEncoder(w).Encode(map[string]string{
	// 		"status":"OK",
	// 		"message":"API Running",
	// 	})
	// })


	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID) 

	addr := "0.0.0.0:" + config.Port

	fmt.Println("Server running on port", config.Port)

	err = http.ListenAndServe(addr,nil)

	if err != nil {
		fmt.Println("Failed running server")
	}
}