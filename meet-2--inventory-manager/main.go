package main

import (
	"fmt"
	"slices"
)

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

const AppName string = "====== Toko Kelontong | INVENTORY MANAGER ======"
var menus = []string{"Tambah Barang Ke Gudang", "Lihat Semua Stock Barang", "Beli Barang", "Keluar"}
var products = []Product{}

func title() {
	
	fmt.Print("\n\n",AppName, "\n\n")
	for index, menu := range menus {
		fmt.Printf("%d. %v\n", index + 1, menu)
	}
}

// func inputInt(q string, errMsg *string) {
// 	var input, _input string

// 	for {
// 		fmt.Print(q)
// 		fmt.Scanln(&_input)

// 		input, err = strconv.Atoi(_input)

// 	}
// }

func task__() {
	var inputItemName string
	var inputMenu, inputPrice, inputStock, inputID, inputCount, inputMoney int

	menuLen := len(menus)
	
	for {
		title()

		fmt.Printf("\n\n\nPilih Menu [1-%d]: ", menuLen)
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			fmt.Printf("Masukan Nama Barang: ")
			fmt.Scanln(&inputItemName)
			
			fmt.Printf("Masukan Harga Barang: ")
			fmt.Scanln(&inputPrice)
			
			fmt.Printf("Masukan Stock Barang: ")
			fmt.Scanln(&inputStock)
			
			// product := Product{
			// 	ID: 	len(products) + 1,
			// 	Name: 	inputItemName,
			// 	Price: 	price,
			// 	Stock: 	stock,
			// }

			products = append(products, Product{
				ID: 	len(products) + 1,
				Name: 	inputItemName,
				Price: 	inputPrice,
				Stock: 	inputStock,
			})

			fmt.Printf("\n[SYSTEM]: Barang berhasil ditambahkan ke gudang.")
		case 2:
			fmt.Print("\n====== Daftar Stock Gudang ======\n")

			for _, product := range products {
				fmt.Printf("ID: %d	| Nama:	%v		| Harga: Rp %d		| Stok: %d", product.ID, product.Name, product.Price, product.Stock)
			}

			fmt.Printf("=====================\nTotal Jenis Barang: %d", len(products) + 1)
		case 3:
			fmt.Printf("Masukan ID Barang yang mau dibeli: ")
			fmt.Scanln(&inputID)

			productIndex := slices.IndexFunc(products, func(p Product) bool {
				return p.ID == inputID
			})
			if productIndex == -1 {
				fmt.Print("[SYSTEM]: Barang dengan ID tersebut tidak ditemukan.")
				continue
			}
			
			fmt.Printf("Masukan jumlah Barang yang mau dibeli: ")
			fmt.Scanln(&inputCount)
			
			fmt.Printf("Masukan uang anda: ")
			fmt.Scanln(&inputMoney)

		case 4:
			fmt.Print("[SYSTEM]: Program selesai. Selamat Beristirahat, Bos.")
			return
		default:
			// 
		}
	}
}

func main() {
	// task__()
	var num int

	fmt.Print("input number: ")
	fmt.Scanln(&num)

	fmt.Printf("%d + 1 = %d", num, num + 1)
}