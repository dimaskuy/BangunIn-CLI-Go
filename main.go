package main

import "fmt"

func menu() int { // dimas - menambahkan fungsi menu
	var input int
	fmt.Println("1. Lihat bangunan")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. Urutkan")
	fmt.Println("5. Cari")
	fmt.Print("Masukkan pilihan (1-5): ")

	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Input harus berupa angka!")
		return 0
	}
	return input
}

func main() {
	input := menu()
	fmt.Print(input)
}
