package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func Cart(in int) {
	category := database.GetCategory()
	cart := database.GetCart()
	var input string
	success := false
	name := ""
	price := 0
	description := ""

	for _, results := range *category {
		search := results.Urut
		if search == in {
			name = results.Name
			price = results.Price
			description = results.Description
			*cart = append(*cart, results.Id)
			success = true
			break
		}
	}

	if !success {
		utils.WrongInput(&input)
		return
	}

	utils.Clear()
	fmt.Printf("\nName   : %s", name)
	fmt.Printf("\nPrice  : %d\n", price)
	fmt.Printf("\nDeskripsi\n")
	fmt.Printf("-------------------------------\n")
	fmt.Println(description)
	fmt.Printf("-------------------------------\n")
	fmt.Printf("\nTekan ENTER untuk melanjutkan...")
	fmt.Scanln(&input)
}
