package utils

import (
	"authenticatiion-flow/database"
	"sync"
)

func AddHistory(wait *sync.WaitGroup, signal *chan string) {
	products := database.GetProducts()
	cart := database.GetCart()
	history := database.GetHistory()
	for _, res := range *cart {
		for i, result := range *products {
			if res.Id == result.Id {
				(*products)[i].Stock = (*products)[i].Stock - res.Qty
			}
		}
		newHistory := database.History{
			Username: UserNameActived(),
			Id:       res.Id,
			Urut:     res.Urut,
			Name:     res.Name,
			Price:    res.Price,
			Qty:      res.Qty,
		}
		*history = append(*history, newHistory)
	}
	*cart = []database.Cart{}
	*signal <- "Transksi Sukses..."
	wait.Done()
}
