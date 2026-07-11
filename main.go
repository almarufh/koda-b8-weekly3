package main

import (
	"authenticatiion-flow/lib"
	"authenticatiion-flow/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup
	signal := make(chan int)
	wait.Add(1)
	go utils.GetProducts(&wait, &signal)
	defer wait.Wait()
	i := 0
	for i < 1 {
		utils.Clear()
		fmt.Println("Loading .")
		time.Sleep(time.Duration(1) * time.Second)
		utils.Clear()
		fmt.Println("Loading . .")
		time.Sleep(time.Duration(1) * time.Second)
		utils.Clear()
		fmt.Println("Loading . . .")
		time.Sleep(time.Duration(1) * time.Second)
		if <-signal == 1 {
			i++
		}
	}
	lib.Menu()
}
