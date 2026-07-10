package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func Dashboard() {
	utils.Clear()
	getActived := database.GetActived()
	actived := *getActived
	var name string
	var indexActived int
	for i := range actived {
		if actived[i].Status == true {
			name = actived[i].Name
			indexActived = i
		}
	}
	var input string
	fmt.Printf("\n--- Welcome to system ---\n\nHello %s\n\n[1] All Products\n[2] Category\n[3] Search Products\n[4] Cart\n[5] Checkout\n\n[0] Logout\n[00] Exit\n\nChoose a menu :   ", name)
	fmt.Scanf("%s", &input)

	switch input {
	case "1":
		return
	case "2":
		ShowCategory()
	case "3":
		return
	case "4":
		return
	case "5":
		return
	case "0":
		actived[indexActived].Status = false
		Menu()
	case "00":
		return
	default:
		Dashboard()
	}
}
