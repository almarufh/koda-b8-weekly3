package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func ShowAllProducts() {
	getCategory := database.GetCategory()
	newCategory := []database.Category{}
	products := database.GetProducts()
	urut := 0

	for _, results := range *products {
		urut += 1
		product := database.Category{
			Urut:        urut,
			Id:          results.Id,
			Name:        results.Name,
			Description: results.Description,
			Price:       results.Price,
			Stock:       results.Stock,
		}
		newCategory = append(newCategory, product)
	}

	*getCategory = newCategory

	for {
		utils.Clear()
		var input int
		qty := len(*getCategory)
		fmt.Printf("--- [ List All Products ] ---\n")
		for i, results := range *getCategory {
			if i <= 8 {
				fmt.Printf("[%d]  %s <----> [Rp%d]\n", i+1, results.Name, results.Price)
			} else {
				fmt.Printf("[%d] %s <----> [Rp%d]\n", i+1, results.Name, results.Price)
			}
		}

		fmt.Printf("\n\n--------------------------\n[0] Back\n\nChose menu (1-%d): ", qty)

		_, err := fmt.Scanln(&input)

		if err != nil {
			utils.WrongInput()
			continue
		}

		if input > 0 && input <= qty {
			Order(input)
			continue
		}

		switch input {
		case 0:
			return
		default:
			utils.WrongInput()
			continue
		}
	}

}
