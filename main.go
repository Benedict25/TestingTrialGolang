package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

// func menu(){
// var menu int
// switch fmt.Scanln(&menu) {
// case 1:
// 	// Menu 1
// case 2:
// 	// Menu 2
// case 3:
// 	// Exit
// 	fmt.Printf("Terima Kasih telah menggunakan layanan kami!")
// 	os.Exit(3)
// default:
// 	// coba lagi
// 	fmt.Printf("Coba lagi")
// }
// 	menu()
// }

// func HitungJumlahMuatan() {
// 	var sisiBox float64 = 0
// 	var sisiGudang float64 = 0
// 	fmt.Printf("Ukuran sisi Box: ")
// 	fmt.Scanln(&sisiBox)
// 	sisiBox = CekSisiBox(sisiBox)

// 	fmt.Printf("Ukuran sisi Gudang: ")
// 	fmt.Scanln(&sisiGudang)
// 	sisiGudang = CekSisiGudang(sisiGudang)

// }

func main() {
	var menu int
	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2}

	for {
		fmt.Print("\nMenu:\n1. Cek Biaya Sewa\n2. Cek Jumlah Muatan\n3. Exit\nMasukkan pilihan menu: ")
		fmt.Scanln(&menu)

		if menu == 3 {
			break
		}

		switch menu {
		case 1:
			var sisiGudang float64

			fmt.Print("\nMasukkan Sisi Gudang: ")
			fmt.Scanln(&sisiGudang)

			sisiGudang = CekSisiGudang(sisiGudang)

			var totalsewa = HitungTotalBayarSewa(sisiGudang)
			fmt.Println("Jumlah yang harus dibayar: ", ac.FormatMoney(totalsewa))

		case 2:
			var sisiGudang float64
			var sisiBox float64

			fmt.Print("\nMasukkan sisi Gudang: ")
			fmt.Scanln(&sisiGudang)

			fmt.Print("\nMasukkan sisi Box: ")
			fmt.Scanln(&sisiBox)

			sisiGudang = CekSisiGudang(sisiGudang)
			sisiBox = CekSisiBox(sisiBox)

			var jumlahmuatan = HitungMuatanGudang(sisiBox, sisiGudang)

			if jumlahmuatan == 0 {
				/*error balikin -> ngisi ulang ato ke menu utama*/
				fmt.Println("Maaf Input anda tidak sesuai/tidak tersedia - Harap Ulangi!")
			} else {
				fmt.Println("Gudang tersebut bisa menampung: ", jumlahmuatan, " box")
			}

		case 3:
			fmt.Print("\nTerima Kasih telah menggunakan layanan kami!\n")

		default:
			fmt.Print("\nInput salah silahkan coba lagi\n")
		}
	}
}
