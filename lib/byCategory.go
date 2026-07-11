package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func ByCategory(category string) {
	getCategory := database.GetCategory()
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
				}
				*getCategory = append(*getCategory, product)
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

				*getCategory = append(*getCategory, product)
			}
		}
	}

	for {
		utils.Clear()
		var input int
		fmt.Printf("\n--- [ Products Category %s ] ---\n\n", category)

		for i, results := range *getCategory {
			if i <= 8 {
				fmt.Printf("[%d] %s\n", i+1, results.Name)
			} else {
				fmt.Printf("[%d]  %s\n", i+1, results.Name)
			}
		}
		fmt.Printf("\n\nChose menu : ")
		fmt.Scanf("%d", &input)
		Cart(input)
		continue
	}

}
