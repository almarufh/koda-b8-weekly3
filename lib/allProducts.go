package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func ShowAllProducts() {
	utils.Clear()
	fmt.Printf("--- [ List All Products ] ---\n")
	products := database.GetProducts()
	for i, results := range *products {
		if i <= 8 {
			fmt.Printf("[%d]  %s\n", i+1, results.Name)
		} else {
			fmt.Printf("[%d] %s\n", i+1, results.Name)
		}
	}
}
