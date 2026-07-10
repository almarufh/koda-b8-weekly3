package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"os"
)

func ShowCategory() {
	utils.Clear()
	var input string
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
		fmt.Println("Masih dalam pengembangan")
	case "2":
		fmt.Println("Masih dalam pengembangan")
		return
	case "3":
		fmt.Println("Masih dalam pengembangan")
		return
	case "4":
		fmt.Println("Masih dalam pengembangan")
		return
	case "0":
		os.Exit(1)
	case "00":
		return
	default:
		ShowCategory()
	}
}
