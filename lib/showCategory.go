package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"os"
)

func ShowCategory() {
	for {
		var input string
		utils.Clear()
		fmt.Printf("\n=======================================")
		fmt.Printf("\nList Menu %s", database.Project)
		fmt.Printf("\n=======================================\n\n")
		fmt.Println("[1] Promotion")
		fmt.Println("[2] Beverage")
		fmt.Println("[3] Chicken")
		fmt.Println("[4] Bowl/Burger")
		fmt.Println("\n\n--------------------------------------")
		fmt.Println("[0] Back")
		fmt.Println("[00] Exit")
		fmt.Printf("\n\n=======================================\n")
		fmt.Printf("Pilih Kategori (1-4) :  ")
		fmt.Scanf("%s", &input)

		switch input {
		case "1":
			ByCategory("Promotion")
		case "2":
			ByCategory("Beverage")
		case "3":
			ByCategory("Chicken")
		case "4":
			ByCategory("Bowl/Burger")
		case "0":
			return
		case "00":
			os.Exit(1)
		default:
			utils.WrongInput()
			continue
		}
	}
}
