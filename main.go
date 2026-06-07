package main

import "fmt"

type Supplier struct {
	namaPerusahaan   string
	lokasi           string
	jenisMaterial    string
	ratingPerforma   int
	detailKontak     string
	riwayatPelayanan string
}

// dimas - mempercantik dengan mengubah warna terminal
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	White  = "\033[37m"
)

func errMsg(str string) string {
	return Red + str + Reset
}
func successMsg(str string) string {
	return Green + str + Reset
}
func warnMsg(str string) string {
	return Yellow + str + Reset
}

func convertRating(rating *int) string {
	switch *rating {
	case 1:
		return "⭐"
	case 2:
		return "⭐⭐"
	case 3:
		return "⭐⭐⭐"
	case 4:
		return "⭐⭐⭐⭐"
	case 5:
		return "⭐⭐⭐⭐⭐"
	default:
		return ""
	}
}

func pause() {
	fmt.Print(warnMsg("\nTekan Enter untuk kembali ke menu!"))
	fmt.Scanln()
	fmt.Scanln()
}

func menu() { // dimas - menambahkan fungsi menu
	var input int
	var suppliers []Supplier

	for {
		fmt.Print(Yellow + "~~ Menu BangunIn ~~" + Reset)
		fmt.Print(Green)

		fmt.Printf(
			"\n%s1.%s Tambah Supplier\n"+
				"%s2.%s Tampilkan Supplier\n"+
				"%s3.%s Cari Supplier\n"+
				"%s4.%s Ubah Rating Performa\n"+
				"%s5.%s Hapus Supplier\n"+
				"%s6.%s Urutkan Supplier (Rating Tertinggi)\n"+
				"%s7.%s Urutkan Supplier (Rating Terendah)\n"+
				"%s8.%s Tampilkan Statistik Aplikasi\n"+
				"%s0.%s Keluar\n"+
				"Masukkan pilihan menu (0-8): ",
			Green, White,
			Green, White,
			Green, White,
			Green, White,
			Green, White,
			Green, White,
			Green, White,
			Green, White,
			Green, White,
		)
		fmt.Print(Reset)
		fmt.Scan(&input)

		if input == 1 {
			suppliers = add(suppliers)
			pause()
		} else if input == 2 {
			show(suppliers)
			pause()
		} else if input == 3 {
			var inputSort int
			fmt.Printf(
				"\n%s1.%s Urut berdasarkan Nama\n"+
					"%s2.%s Urut berdasarkan Lokasi\n"+
					"Masukkan pilihan (1/2): ",
				Green, White,
				Green, White,
			)
			fmt.Scan(&inputSort)
			if inputSort == 1 {
				searchByName(suppliers)
				pause()
			} else if inputSort == 2 {
				searchByLocation(suppliers)
				pause()
			}
		} else if input == 4 {
			suppliers = editRating(suppliers)
			pause()
		} else if input == 5 {
			suppliers = delete(suppliers)
			pause()
		} else if input == 6 {
			suppliers = sortRatingAsc(suppliers)
			pause()
		} else if input == 7 {
			suppliers = sortRatingDesc(suppliers)
			pause()
		} else if input == 8 {
			showStats(suppliers)
			pause()
		} else if input == 0 {
			fmt.Println(errMsg("Berhasil keluar dari aplikasi!"))
			break
		} else {
			fmt.Println(errMsg("Pilihan tidak valid, coba lagi!"))
		}
	}
}

// wira - mengubah fungsi agar bisa input lebih dari 1 data sekaligus
func add(list []Supplier) []Supplier {
	fmt.Println(warnMsg(">>> Tambah Supplier <<<"))

	var jumlah int
	fmt.Print("Banyak data supplier: ")
	fmt.Scan(&jumlah)

	for i := 1; i <= jumlah; i++ {
		var newData Supplier
		fmt.Printf("\n*Input Data Supplier ke-%d ---\n", i)

		fmt.Print("Masukkan Nama Perusahaan: ")
		fmt.Scan(&newData.namaPerusahaan)

		fmt.Print("Lokasi: ")
		fmt.Scan(&newData.lokasi)

		fmt.Print("Jenis Material: ")
		fmt.Scan(&newData.jenisMaterial)

		for {
			fmt.Print("Rating Performa (1-5): ")
			fmt.Scan(&newData.ratingPerforma)

			if newData.ratingPerforma >= 1 && newData.ratingPerforma <= 5 {
				break
			}
			fmt.Println(errMsg("Rating harus dalam rentang 1 hingga 10, coba lagi."))
		}

		fmt.Print("Detail Kontak: ")
		fmt.Scan(&newData.detailKontak)

		fmt.Print("Masukkan Riwayat/Deskripsi Pelayanan: ")
		fmt.Scan(&newData.riwayatPelayanan)

		// slice arr
		list = append(list, newData)
		fmt.Printf("Data supplier ke-%d berhasil disimpan!\n", i)
	}

	fmt.Println(successMsg("\n Semua data baru berhasil tersimpan"))
	fmt.Println("-----------------------------------------------------")
	return list
}

// wira - menambahkan function tambahkan supplier
func show(list []Supplier) {
	fmt.Println(warnMsg(">>> Tampilkan Supplier <<<"))

	if len(list) == 0 {
		fmt.Println(errMsg("Belum ada data supplier yang tersimpan!"))
		fmt.Println("-----------------------------------")
		return
	} else {
		for _, data := range list {
			fmt.Println("-----------------------------------")
			fmt.Println(successMsg("Nama Perusahaan   :"), data.namaPerusahaan)
			fmt.Println(successMsg("Lokasi            :"), data.lokasi)
			fmt.Println(successMsg("Jenis Material    :"), data.jenisMaterial)
			fmt.Println(successMsg("Rating Performa   :"), convertRating(&data.ratingPerforma))
			fmt.Println(successMsg("Detail Kontak     :"), data.detailKontak)
			fmt.Println(successMsg("Riwayat Pelayanan :"), data.riwayatPelayanan)
		}
		fmt.Println("-----------------------------------")
	}
}

// BINARY SEARCH
// dimas - menambahkan function cari supplier
func searchByName(list []Supplier) {
	fmt.Println(warnMsg(">>> Cari Supplier <<<"))

	var name string
	fmt.Print("Masukkan Nama Perusahaan yang dicari: ")
	fmt.Scan(&name)

	found := false
	for _, data := range list {
		if data.namaPerusahaan == name {
			fmt.Println("-----------------------------------")
			fmt.Println("Data Ditemukan!")
			fmt.Println(warnMsg("Nama Perusahaan   :"), data.namaPerusahaan)
			fmt.Println(successMsg("Lokasi            :"), data.lokasi)
			fmt.Println(successMsg("Jenis Material    :"), data.jenisMaterial)
			fmt.Println(successMsg("Rating Performa   :"), data.ratingPerforma)
			fmt.Println(successMsg("Detail Kontak     :"), data.detailKontak)
			fmt.Println(successMsg("Riwayat Pelayanan :"), data.riwayatPelayanan)
			fmt.Println("-----------------------------------")
			found = true
		}
	}

	if !found {
		fmt.Println(errMsg("Maaf, supplier dengan nama tersebut tidak ditemukan."))
	}
	fmt.Println("-----------------------------------")
}

// SEQUENTAL SEARCH
// wira - menambahkan function cari supplier berdasarkan lokasi (Sequential Search)
func searchByLocation(list []Supplier) {
	fmt.Println(warnMsg(">>> Cari Supplier Berdasarkan Lokasi <<<"))

	if len(list) == 0 {
		fmt.Println(errMsg("Belum ada data supplier yang tersimpan!"))
		return
	}

	var cariLokasi string
	fmt.Print("Masukkan lokasi: ")
	fmt.Scan(&cariLokasi)

	found := false

	for _, data := range list {
		if data.lokasi == cariLokasi {
			if !found {
				fmt.Println("\nData Ditemukan:")
			}

			fmt.Println("-----------------------------------")
			fmt.Println(successMsg("Nama Perusahaan   :"), data.namaPerusahaan)
			fmt.Println(warnMsg("Lokasi            :"), data.lokasi)
			fmt.Println(successMsg("Jenis Material    :"), data.jenisMaterial)
			fmt.Println(successMsg("Rating Performa   :"), convertRating(&data.ratingPerforma))
			fmt.Println(successMsg("Detail Kontak     :"), data.detailKontak)
			fmt.Println(successMsg("Riwayat Pelayanan :"), data.riwayatPelayanan)
			fmt.Println("-----------------------------------")

			found = true
		}
	}

	if !found {
		fmt.Println(errMsg("Supplier pada lokasi tersebut tidak ditemukan."))
	}

	fmt.Println("-----------------------------------")
}

// dimas - menambahkan function ubah rating
func editRating(list []Supplier) []Supplier {
	fmt.Println(warnMsg(">>> Edit Rating Supplier <<<"))

	var target string
	fmt.Print("Masukkan Nama Perusahaan yang akan diubah: ")
	fmt.Scan(&target)

	found := false
	for i, data := range list {
		if data.namaPerusahaan == target {
			var ratingBaru int
			for {
				fmt.Print("Masukkan Rating Baru (1-5): ")
				fmt.Scan(&ratingBaru)

				if ratingBaru >= 1 && ratingBaru <= 5 {
					break
				}
			}

			list[i].ratingPerforma = ratingBaru
			fmt.Println("Rating berhasil diperbarui!")
			found = true
		}
	}

	if !found {
		fmt.Println(errMsg("Maaf, supplier dengan nama tersebut tidak ditemukan."))
	}
	fmt.Println("--------------------")
	return list
}

// wira - menambahkan hapus supplier
func delete(list []Supplier) []Supplier {
	fmt.Println(warnMsg(">>> Hapus Supplier <<<"))

	var target string
	fmt.Print("Masukkan Nama Perusahaan yang akan dihapus: ")
	fmt.Scan(&target)

	found := false
	for i, data := range list {
		if data.namaPerusahaan == target {
			list = append(list[:i], list[i+1:]...)
			fmt.Println(successMsg("Supplier berhasil dihapus!"))
			found = true
			break
		}
	}

	if !found {
		fmt.Println(errMsg("Maaf, supplier dengan nama tersebut tidak ditemukan."))
	}
	return list
}

// SELECTION SORT
// wira - menambahkan alur selection
func sortRatingAsc(list []Supplier) []Supplier {
	fmt.Println(warnMsg(">>> Urutkan Supplier (Rating Tertinggi) <<<"))
	n := len(list)
	if n == 0 {
		fmt.Println(errMsg("Belum ada data supplier yang bisa diurutkan!"))
		return list
	}

	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if list[j].ratingPerforma > list[idxMax].ratingPerforma {
				idxMax = j
			}
		}
		list[i], list[idxMax] = list[idxMax], list[i]
	}

	fmt.Println(successMsg("Data berhasil diurutkan berdasarkan rating tertinggi!"))
	show(list)
	return list
}

// INSERTION SORT
// dimas - menambahkan alur insertion
func sortRatingDesc(list []Supplier) []Supplier {
	fmt.Println(warnMsg(">>> Urutkan Supplier (Rating Terendah) <<<"))
	n := len(list)
	if n == 0 {
		fmt.Println(errMsg("Belum ada data supplier yang bisa diurutkan!"))
		fmt.Println("--------------------")
		return list
	}

	for i := 1; i < n; i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && list[j].ratingPerforma > key.ratingPerforma {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = key
	}

	fmt.Println(successMsg("Data berhasil diurutkan berdasarkan rating terendah"))
	fmt.Println("--------------------")
	show(list)
	return list
}

// wira - menambahkan untuk menampilkan statistik total dan rata rata
func showStats(list []Supplier) {
	fmt.Println(warnMsg(">>> Statistik Aplikasi <<<"))
	n := len(list)
	if n == 0 {
		fmt.Println(errMsg("Belum cukup data supplier untuk statistik!"))
		fmt.Println("--------------------")
		return
	}

	var totalRating int
	for _, data := range list {
		totalRating += data.ratingPerforma
	}
	avgRating := totalRating / n

	fmt.Println("--- Jumlah Supplier Per Wilayah ---")
	var wilayahDihitung []string

	for _, dataOut := range list {
		isCount := false
		for _, w := range wilayahDihitung {
			if w == dataOut.lokasi {
				isCount = true
				break
			}
		}

		if !isCount {
			hitungWilayah := 0
			for _, dataDalam := range list {
				if dataDalam.lokasi == dataOut.lokasi {
					hitungWilayah++
				}
			}
			fmt.Printf("- Wilayah %s: %d Supplier\n", dataOut.lokasi, hitungWilayah)
			wilayahDihitung = append(wilayahDihitung, dataOut.lokasi)
		}
	}

	fmt.Println("-----------------------------------")
	fmt.Printf("Rata-rata Skor Kepuasan Mitra: %.2f\n", avgRating)
	fmt.Println("-----------------------------------")
}

func main() {
	fmt.Print(successMsg(`
██████╗  █████╗ ███╗   ██╗ ██████╗ ██╗   ██╗███╗   ██╗██╗███╗   ██╗
██╔══██╗██╔══██╗████╗  ██║██╔════╝ ██║   ██║████╗  ██║██║████╗  ██║
██████╔╝███████║██╔██╗ ██║██║  ███╗██║   ██║██╔██╗ ██║██║██╔██╗ ██║
██╔══██╗██╔══██║██║╚██╗██║██║   ██║██║   ██║██║╚██╗██║██║██║╚██╗██║
██████╔╝██║  ██║██║ ╚████║╚██████╔╝╚██████╔╝██║ ╚████║██║██║ ╚████║
╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═╝╚═╝  ╚═══╝
	`))
	fmt.Println(Blue + "made with ❤️  by @dimasramadhans && @jalawirabahari // IF-05-02" + Reset)

	menu()
}
