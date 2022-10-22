package main

import (
	"math"
)

// Program "Only-Box" - Storage System
// Pilih ukuran gudang, pilih ukuran box (pasti selalu sama - fixed size)

func HitungVolume(sisi float64) float64 {

	if sisi <= 0 {
		//fmt.Println("Error! Check Input - Input tidak boleh <= 0")
		return 0
	} else {
		var volume float64 = math.Pow(sisi, 3)
		//fmt.Println("Volumenya adalah: ", volume)
		return volume
	}

}

func HitungMuatanGudang(sisiBox float64, sisiGudang float64) float64 {

	var volumeBox float64 = HitungVolume(sisiBox)
	var volumeGudang float64 = HitungVolume(sisiGudang)

	if volumeBox == 0 || volumeGudang == 0 {
		//fmt.Println("Error! Volume Tidak Dapat Diproses")
		return 0
	} else if volumeBox > volumeGudang {
		//fmt.Println("Error! Volume Box Lebih Besar Dari Gudang")
		return 0
	} else {
		var muatan float64 = math.Floor(volumeGudang / volumeBox)
		//fmt.Println("Total barang yang dapat ditampung adalah: ", muatan)
		return muatan
	}

}

func HitungTotalBayarSewa(sisiGudang float64) float64 {

	var volumeGudang float64 = HitungVolume(sisiGudang)

	// Harga Gudang = 100.000 / m^3
	var totalBayar float64 = volumeGudang * 100000

	// Diskon 10% Jika Total Bayar > 500.000.000
	if totalBayar > 500000000 {
		return totalBayar * 0.90
	} else {
		return totalBayar
	}

}
