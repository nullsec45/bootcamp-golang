package main

import (
	"fmt"
	"net/http"
	"strings"
	"kasir-api/database"
	"kasir-api/handlers"
	"os"
	"log"
	"kasir-api/repositories"
	"kasir-api/services"
	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
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


	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	defer db.Close()

	categoryRepo := repositories.NewCategoryRepository(db)
	CategoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(CategoryService)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo, categoryRepo)
	productHandler := handlers.NewProductHandler(productService)

	

	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID) 

	addr := "0.0.0.0:" + config.Port

	fmt.Println("Server running on port", config.Port)

	err = http.ListenAndServe(addr,nil)

	if err != nil {
		fmt.Println("Failed running server")
	}
}