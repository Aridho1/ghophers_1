package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

const AppName string = "====== [ Toko Kelontong | INVENTORY MANAGER ] ======"
const line = "-----------------------"
var reader = bufio.NewReader(os.Stdin)
var menus = []string{"Tambah Barang Ke Gudang", "Lihat Semua Stock Barang", "Beli Barang", "Keluar"}
var products = []Product{}

func title(loopCount int) {

	newLine := "\n\n"
	if loopCount != 0 {
		newLine = "\n\n\n\n"
	}
	
	fmt.Printf("%v%v\n\n", newLine, AppName)
	// fmt.Printf("\n\n\n\n",AppName, "\n\n")
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

func getListProduct(title string) bool {
	productsLen := len(products)

	if productsLen == 0 {
		fmt.Printf("[SYSTEM]: Kita belum memiliki produk apapun di gudang\n!")
		return false
	}

    fmt.Printf("\n\n====== %s ======\n", title)
	
	for i, product := range products {

		barrier := "---------------------------------"
		if i + 1 == productsLen {
			barrier = "================================="
		}

		fmt.Printf("ID: %v	| Nama	: %v	\n	| Harga	: %v	\n	| Stok	: %v\n%v\n", product.ID, product.Name, product.Price, product.Stock, barrier)
		// fmt.Printf("ID: %v | Nama: %v=|=Harga: Rp %v | Stok: %v\n", product.ID, product.Name, product.Price, product.Stock)
	}
	
	// fmt.Print("\n")
	return true
}

func task__() {
	// var inputItemName string
	var inputMenu, inputPrice, inputStock, inputID, inputCount, inputMoney int

	menuLen := len(menus)
	
	for loopCount := 0; loopCount >= 0; loopCount++ {
		title(loopCount)

		if loopCount != 0 {
			time.Sleep(2 * time.Second)	
		}

		fmt.Printf("\nPilih Menu [1-%d]: ", menuLen)
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			fmt.Printf("Masukan Nama Barang: ")
			// fmt.Scanln(&inputItemName)
			inputItemName, _ := reader.ReadString('\n')
			inputItemName = strings.TrimSpace(inputItemName)
				
			fmt.Printf("Masukan Harga Barang: ")
			fmt.Scanln(&inputPrice)
			
			fmt.Printf("Masukan Stock Barang: ")
			fmt.Scanln(&inputStock)

			products = append(products, Product{
				ID: 	len(products) + 1,
				Name: 	inputItemName,
				Price: 	inputPrice,
				Stock: 	inputStock,
			})

			fmt.Printf("\n[SYSTEM]: Barang berhasil ditambahkan ke gudang.\n")
		case 2:
			getListProduct("Daftar Stock Gudang")
			fmt.Printf("\nTotal Jenis Barang: %d\n", len(products))
		case 3:
			getListProduct("Daftar Produk")
			fmt.Printf("\nMasukan ID Barang yang mau dibeli: ")
			fmt.Scanln(&inputID)

			productIndex := slices.IndexFunc(products, func(p Product) bool {
				return p.ID == inputID
			})
			if productIndex == -1 {
				fmt.Print("[SYSTEM]: Barang dengan ID tersebut tidak ditemukan.\n")
				continue
			}

			product := &products[productIndex]
			
			fmt.Printf("Masukan jumlah Barang yang mau dibeli: ")
			fmt.Scanln(&inputCount)

			price := inputCount * product.Price
			
			fmt.Printf("[SYSTEM]: Total Harga: Rp %d\n", price)
			fmt.Printf("Masukan uang anda: ")
			fmt.Scanln(&inputMoney)

			changeMoney := inputMoney - price

			if changeMoney < 0 {
				fmt.Printf("[SYSTEM]: Uang anda kurang %d !!!\n", changeMoney * -1)
				continue
			}

			product.Stock -= inputCount

			fmt.Printf("\n[SYSTEM]: Transaksi berhasil!\n")
			fmt.Printf("[SYSTEM]: Kembalian: %d\n", changeMoney)
			fmt.Printf("[SYSTEM]: Stok %v sekarang: %d pcs\n", product.Name, product.Stock)
			// fmt.Printf("%#+ lalu %#v\n\n", product, products)

		case 4:
			if loopCount == 0 {
				fmt.Printf("\n[SYSTEM]: 	Sedih... :(\n		Kamu bahkan belum melakukan aksi apapun :(\n\n")
			} else {
				fmt.Print("[SYSTEM]: Program selesai. Selamat Beristirahat, Bos.")
			}
			return
		}
	}
}

func main() {
	task__()
}