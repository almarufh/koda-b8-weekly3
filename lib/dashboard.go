package lib

import (
	"authenticatiion-flow/utils"
	"fmt"
	"os"
)

func Dashboard() {
	for {
		utils.Clear()
		var input string
		fmt.Printf("\n--- Welcome to system ---\n\nHello %s\n\n[1] All Products\n[2] Category\n[3] Search Products\n[4] Cart\n[5] Checkout\n\n[0] Logout\n[00] Exit\n\nChoose a menu :   ", utils.NameActived())
		fmt.Scanf("%s", &input)

		switch input {
		case "1":
			ShowAllProducts()
			return
		case "2":
			ShowCategory()
			return
		case "3":
			return
		case "4":
			return
		case "5":
			return
		case "0":
			check := utils.Logout()
			if check {
				Menu()
			}
		case "00":
			os.Exit(1)
		default:
			utils.WrongInput()
			continue
		}
	}
}
