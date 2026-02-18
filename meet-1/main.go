package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("hello world")
	
	// // txt, err := fmt.Scanln("test: ")
	// var txt string
	
	// fmt.Printf("Test: ")
	// fmt.Scanln(&txt)

	// fmt.Printf("hasil: %s", txt)
	
	var inputPrice, inputMoney string
	// var price, money int
	// var err Error

	fmt.Printf("\n\n========================\n\n")
	fmt.Printf("Harga Barang: ")
	fmt.Scanln(&inputPrice)

	price, err := strconv.Atoi(inputPrice)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			fmt.Printf("Error, inputnya harus int, %s itu bukan int :D", inputPrice)
			} else {
				fmt.Println("Error: ", err)
		}

		main()
		return
	}

	fmt.Printf("Uang yang dimiliki: ")
	fmt.Scanln(&inputMoney)
	money, err := strconv.Atoi(inputMoney)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			fmt.Printf("Error, inputnya harus int, %s itu bukan int :D", inputMoney)
			} else {
				fmt.Println("Error: ", err)
		}

		main()
		return
	}

	var sisa int = money - price

	if sisa < 0 {
		fmt.Println("[SYSTEM]: Transaksi ditolak! Uang kurang ", price * -1)
		main()
		return
	}

	fmt.Println("[SYSTEM]: Transaksi berhasil. Kembalian anda ", sisa)

	fmt.Printf("\n\n========================\n\n")
}