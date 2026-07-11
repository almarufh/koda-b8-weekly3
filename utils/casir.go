package utils

import (
	"authenticatiion-flow/database"
	"sync"
)

func Casir(key int, id string, qty int, price int, signal *chan int, total *int, wait *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	products := database.GetProducts()
	*total += price * qty
	*signal <- key
	for i := range *products {
		if (*products)[i].Id == id {
			(*products)[i].Stock = (*products)[i].Stock - qty
		}
	}
	mut.Unlock()
	wait.Done()
}
