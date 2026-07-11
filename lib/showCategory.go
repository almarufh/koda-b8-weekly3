package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"os"
)

func ShowCategory() {
	var input string
	for {
		utils.Clear()
		fmt.Printf("\n=======================================")
		fmt.Printf("\nList Menu %s", database.Project)
		fmt.Printf("\n=======================================\n\n")
		fmt.Println("[1] Promotion")
		fmt.Println("[2] Beverage")
		fmt.Println("[3] Chicken")
		fmt.Println("[4] Bowl/Burger")
		fmt.Println("\n\n--------------------------------------")
		fmt.Println("[0] Exit")
		fmt.Println("[00] Back")
		fmt.Printf("\n\n=======================================\n")
		fmt.Printf("Pilih Kategori (1-4) :  ")
		fmt.Scanf("%s", &input)

		switch input {
		case "1":
			ByCategory("Promotion")
			return
		case "2":
			ByCategory("Beverage")
			return
		case "3":
			ByCategory("Chicken")
			return
		case "4":
			ByCategory("Bowl/Burger")
			return
		case "0":
			os.Exit(1)
			return
		case "00":
			return
		default:
			utils.WrongInput()
			continue
		}
	}
}
