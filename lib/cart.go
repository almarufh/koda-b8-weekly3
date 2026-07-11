package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func Cart(id string, name string, price int, input int) {
	cart := database.GetCart()
	found := false

	for _, result := range *cart {
		if result.Id == id {
			found = true
		}
	}

	if found == false {
		newOrder := database.Cart{
			Id:    id,
			Urut:  len(*cart) + 1,
			Name:  name,
			Price: price,
			Qty:   input,
		}
		*cart = append(*cart, newOrder)
	} else {
		for i, result := range *cart {
			if result.Id == id {
				(*cart)[i].Qty = result.Qty + input
				break
			}
		}
	}

	for {
		utils.Clear()
		var input string
		fmt.Printf("\n--- [ LIST PESANAN ] ---\n\n")
		for _, res := range *cart {
			fmt.Printf("%d. %s qty(x%d) Rp%d\n", res.Urut, res.Name, res.Qty, res.Price*res.Qty)
		}
		fmt.Printf("\n\n-----------------------\n")
		fmt.Printf("[1] Checkout\n[2] Tambah Pesanan\n\n[0] Logout")
		fmt.Printf("Chose a menu : ")
		fmt.Scanf("%s", &input)
		switch input {
		case "1":
			Checkout()
			continue
		case "2":
			Dashboard()
			return
		case "0":
			check := utils.Logout()
			if check {
				Menu()
			}
			return
		default:
			utils.WrongInput()
			continue
		}
	}
}
