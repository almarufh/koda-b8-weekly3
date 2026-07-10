package utils

import (
	"authenticatiion-flow/database"
	"encoding/json"
	"fmt"
	"os"
)

func GetProducts() *[]database.Products {
	products := database.GetProducts()
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			os.Stdout.Close()
		}
	}()

	file, err := os.ReadFile("./database/products.json")
	if err != nil {
		panic("Mohon maaf list menu tidak ditemukan")
	}
	errMar := json.Unmarshal([]byte(string(file)), products)
	if errMar != nil {
		panic(errMar.Error())
	}

	return products
}
