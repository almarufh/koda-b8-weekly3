package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"os"
)

func Dashboard() {
	for {
		utils.Clear()
		cart := database.GetCart()
		var input string
		fmt.Printf("Hello %s!\nWellcome to %s\n\n[1] All Products\n[2] Category\n[3] Search Products\n", utils.NameActived(), database.Project)
		fmt.Printf("[4] History\n")
		if len(*cart) > 0 {
			for _, val := range *cart {
				if val.Username == utils.UserNameActived() {
					fmt.Printf("[5] Cart\n[6] Checkout\n")
				}
			}
		}
		fmt.Printf("\n[0] Logout\n[00] Exit\n\nChoose a menu :   ")
		fmt.Scanf("%s", &input)

		switch input {
		case "1":
			ShowAllProducts()
			continue
		case "2":
			ShowCategory()
			continue
		case "3":
			Search()
			continue
		case "4":
			History()
			continue
		case "5":
			Cart()
			continue
		case "6":
			Checkout()
			continue
		case "0":
			check := utils.Logout()
			if check {
				Menu()
				return
			}
		case "00":
			os.Exit(1)
		default:
			utils.WrongInput()
			continue
		}
	}
}
