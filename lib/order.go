package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func Order(in int) {
	category := database.GetCategory()
	success := false
	name := ""
	price := 0
	description := ""
	stock := 0
	id := ""

	if success == false {
		for _, results := range *category {
			search := results.Urut
			if search == in {
				name = results.Name
				price = results.Price
				description = results.Description
				stock = results.Stock
				id = results.Id
				success = true
				break
			}
		}
	}

	if !success {
		utils.WrongInput()
		return
	}

	for {
		utils.Clear()
		var input int
		fmt.Printf("\nName   : %s", name)
		fmt.Printf("\nPrice  : %d\n", price)
		fmt.Printf("\nDeskripsi\n")
		fmt.Printf("-------------------------------\n")
		fmt.Println(description)
		fmt.Printf("-------------------------------\n\n")
		fmt.Printf("[0] Back\n")
		fmt.Printf("\nKonfirmasi Jumlah pesanan (1-%d) : ", stock)
		_, err := fmt.Scanln(&input)

		if err != nil {
			utils.WrongInput()
			continue
		}

		if input <= stock && input > 0 {
			Cart(id, name, price, input)
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
