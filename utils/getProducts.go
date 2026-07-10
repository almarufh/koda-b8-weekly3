package utils

import (
	"authenticatiion-flow/database"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

func GetProducts(wg *sync.WaitGroup, signal *chan int) {
	products := database.GetProducts()
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			os.Stdout.Close()
		}
	}()

	file, err := os.ReadFile("./database/products.json")
	if err != nil {
		panic("Mohon maaf list menu tidak ditemukan")
	}
	err = json.Unmarshal([]byte(string(file)), products)
	if err != nil {
		panic("Data products gagal dimuat")
	}
	*signal <- 1
	wg.Done()
}
