package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func ByCategory(category string) {
	getCategory := database.GetCategory()
	newCategory := []database.Category{}
	products := database.GetProducts()
	urut := 0

	if category == "Promotion" {
		for _, results := range *products {
			if results.IsPromo == true {
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
		}
	} else {
		for _, results := range *products {
			if results.Category == category {
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
		}
	}

	*getCategory = newCategory

	for {
		utils.Clear()
		var input int
		fmt.Printf("\n--- [ Products Category %s ] ---\n\n", category)
		qty := len(*getCategory)

		for i, results := range *getCategory {
			if i <= 8 {
				fmt.Printf("[%d] %s\n", i+1, results.Name)
			} else {
				fmt.Printf("[%d]  %s\n", i+1, results.Name)
			}
		}
		fmt.Printf("\n\n--------------------------\n[0] Back\n\nChose menu (1-%d): ", qty)
		// fmt.Scanf("%d", &input)

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
