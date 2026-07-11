package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"sync"
	"time"
)

func Checkout() {
	wait := sync.WaitGroup{}
	mut := sync.Mutex{}
	cart := database.GetCart()

	total := 0
	count := 0
	signal := make(chan int)
	for i, res := range *cart {
		wait.Add(1)
		count += i + 1
		go utils.Casir(i+1, res.Id, res.Qty, res.Price, &signal, &total, &wait, &mut)
	}

	i := 0

	for i != len(*cart) {
		utils.Clear()
		fmt.Printf("--- [ STRUK PEMBELIAN ] ---\n\n")
		for _, res := range *cart {
			fmt.Printf("%d. %s qty(x%d) Rp%d\n", res.Urut, res.Name, res.Qty, res.Price*res.Qty)
		}
		fmt.Printf("\n------------------------")
		fmt.Printf("\nStatus : Belum Bayar")
		fmt.Println("\nTotal : Calculating .")
		time.Sleep(time.Duration(1) * time.Second)
		utils.Clear()
		fmt.Printf("--- [ STRUK PEMBELIAN ] ---\n\n")
		for _, res := range *cart {
			fmt.Printf("%d. %s qty(x%d) Rp%d\n", res.Urut, res.Name, res.Qty, res.Price*res.Qty)
		}
		fmt.Printf("\n------------------------")
		fmt.Printf("\nStatus : Belum Bayar")
		fmt.Println("\nTotal : Calculating . .")
		time.Sleep(time.Duration(1) * time.Second)
		utils.Clear()
		fmt.Printf("--- [ STRUK PEMBELIAN ] ---\n\n")
		for _, res := range *cart {
			fmt.Printf("%d. %s qty(x%d) Rp%d\n", res.Urut, res.Name, res.Qty, res.Price*res.Qty)
		}
		fmt.Printf("\n------------------------")
		fmt.Printf("\nStatus : Belum Bayar")
		fmt.Println("\nTotal : Calculating . . .")
		time.Sleep(time.Duration(1) * time.Second)
		utils.Clear()
		fmt.Printf("--- [ STRUK PEMBELIAN ] ---\n\n")
		for _, res := range *cart {
			fmt.Printf("%d. %s qty(x%d) Rp%d\n", res.Urut, res.Name, res.Qty, res.Price*res.Qty)
		}
		fmt.Printf("\n------------------------")
		fmt.Printf("\nStatus : Belum Bayar")
		fmt.Println("\nTotal : Calculating . . .")
		time.Sleep(time.Duration(1) * time.Second)
		if <-signal == len(*cart) {
			i = len(*cart)
		}
	}

	wait.Wait()

	for {
		var input int
		utils.Clear()
		fmt.Printf("--- [ STRUK PEMBELIAN ] ---\n\n")
		for _, res := range *cart {
			fmt.Printf("%d. %s qty(x%d) Rp%d\n", res.Urut, res.Name, res.Qty, res.Price*res.Qty)
		}
		fmt.Printf("\n------------------------")
		fmt.Printf("\nStatus : Belum Bayar")
		fmt.Printf("\nTotal : Rp%d\n\n", total)
		fmt.Printf("Total Bayar : ")
		_, err := fmt.Scanln(&input)

		if err != nil {
			utils.WrongInput()
			continue
		} else {
			if input < total {
				fmt.Printf("\n\nPembayaran kurang sebesar Rp%d", total-input)
				fmt.Printf("\n\nPress ENTER to input ulang")
				fmt.Scanln()
				continue
			}

			if input >= total {
				utils.Clear()
				fmt.Printf("--- [ STRUK PEMBELIAN ] ---\n\n")
				for _, res := range *cart {
					fmt.Printf("%d. %s qty(x%d) Rp%d\n", res.Urut, res.Name, res.Qty, res.Price*res.Qty)
				}
				fmt.Printf("\n------------------------")
				fmt.Printf("\nStatus : LUNAS")
				fmt.Printf("\nTotal : Rp%d\n\n", total)

				fmt.Println("Pembayaran : Rp", input)
				fmt.Println("Kembalian : Rp", input-total)
				fmt.Printf("\nTerimakasih %s ....", utils.NameActived())
				fmt.Printf("\n\nPres ENTER back to HOME  ")
				fmt.Scanln()
				Dashboard()
			}
		}

	}
}
