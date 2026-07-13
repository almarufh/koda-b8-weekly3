package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func History() {
	utils.Clear()
	history := database.GetHistory()
	fmt.Printf("Hello %s...\n\n--- [ HISTORY PESANAN ] ---\n\n", utils.NameActived())
	for _, res := range *history {
		if res.Username == utils.UserNameActived() {
			fmt.Printf("Hello, %s\n\n", utils.NameActived())
			fmt.Printf("Press ENTER to back....")
			fmt.Printf("\n")
			return
		}
	}
	for i, res := range *history {
		if utils.UserNameActived() == res.Username {
			fmt.Printf("%d. %s qty(%d) Rp%d\n", i+1, res.Name, res.Qty, res.Qty*res.Price)
		}
	}

	fmt.Printf("\n\nPress ENTER to back....")
	fmt.Scanf("\n")
}
