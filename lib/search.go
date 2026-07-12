package lib

import (
	"authenticatiion-flow/utils"
	"fmt"
)

func Search() {
	utils.Clear()
	var search string
	fmt.Printf("\nSearch what you need products :  ")
	fmt.Scanf("%s", &search)
	SearchProduct(search)
}
