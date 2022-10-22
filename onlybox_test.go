package main

import "testing"

var ( // Var - General
	volume     float64
	muatan     float64
	totalBayar float64

	sisi       float64
	sisiBox    float64
	sisiGudang float64
)

// 1. Volume - Sisi Dibawah 0

var (
	volumeSeharusnyaSisiDibawah0 float64 = 0
)

func TestHitungVolumeSisiDibawah0(t *testing.T) {

	sisi = -2.5
	volume = HitungVolume(sisi)

	if volume == volumeSeharusnyaSisiDibawah0 {
		t.Log("/// Volume: ", volume, " Expected: ", volumeSeharusnyaSisiDibawah0, " ///")
	} else {
		t.Error("/// SALAH! Volume: ", volume, " Expected: ", volumeSeharusnyaSisiDibawah0, " ///")
	}

}

// 2. Volume - Sisi = 0

var (
	volumeSeharusnyaSisi0 float64 = 0
)

func TestHitungVolumeSisi0(t *testing.T) {

	sisi = 0
	volume = HitungVolume(sisi)

	if volume == volumeSeharusnyaSisiDibawah0 {
		t.Log("/// Volume: ", volume, " Expected: ", volumeSeharusnyaSisi0, " ///")
	} else {
		t.Error("/// SALAH! Volume: ", volume, " Expected: ", volumeSeharusnyaSisi0, " ///")
	}

}

// 3. Volume - Sukses - Sisi Diatas 0 (2.5)

var (
	volumeSeharusnyaSisiDiatas0 float64 = 15.625
)

func TestHitungVolumeSisiDiatas0(t *testing.T) {

	sisi = 2.5
	volume = HitungVolume(sisi)

	if volume == volumeSeharusnyaSisiDiatas0 {
		t.Log("/// Volume : ", volume, " Expected: ", volumeSeharusnyaSisiDiatas0, " ///")
	} else {
		t.Error("/// SALAH! Volume: ", volume, " Expected: ", volumeSeharusnyaSisiDiatas0, " ///")
	}

}

// 4. Muatan Gudang - volumeBox 0
var (
	muatanSeharusnyaJikaVolumeBox0 float64 = 0
)

func TestHitungMuatanVolumeBox0(t *testing.T) {
	sisiBox = 0
	// sisiBox = -2
	sisiGudang = 5
	muatan = HitungMuatanGudang(sisiBox, sisiGudang)

	if muatan == muatanSeharusnyaJikaVolumeBox0 {
		t.Log("/// Muatan : ", muatan, " Expected: ", muatanSeharusnyaJikaVolumeBox0, " ///")
	} else {
		t.Error("/// SALAH! Muatan: ", muatan, " Expected: ", muatanSeharusnyaJikaVolumeBox0, " ///")
	}
}

// 5. Muatan Gudang - volumeGudang 0

var (
	muatanSeharusnyaJikaVolumeGudang0 float64 = 0
)

func TestHitungMuatanVolumeGudang0(t *testing.T) {
	sisiBox = 1
	sisiGudang = 0
	muatan = HitungMuatanGudang(sisiBox, sisiGudang)

	if muatan == muatanSeharusnyaJikaVolumeGudang0 {
		t.Log("/// Muatan : ", muatan, " Expected: ", muatanSeharusnyaJikaVolumeGudang0, " ///")
	} else {
		t.Error("/// SALAH! Muatan: ", muatan, " Expected: ", muatanSeharusnyaJikaVolumeGudang0, " ///")
	}

}

// 6. Muatan Gudang - volumeBox 0 & volumeGudang 0

var (
	muatanSeharusnyaJikaVolumeBoxDanVolumeGudang0 float64 = 0
)

func TestHitungMuatanVolumeBoxDanVolumeGudang0(t *testing.T) {
	sisiBox = 0
	sisiGudang = 0
	muatan = HitungMuatanGudang(sisiBox, sisiGudang)

	if muatan == muatanSeharusnyaJikaVolumeBoxDanVolumeGudang0 {
		t.Log("/// Muatan : ", muatan, " Expected: ", muatanSeharusnyaJikaVolumeBoxDanVolumeGudang0, " ///")
	} else {
		t.Error("/// SALAH! Muatan: ", muatan, " Expected: ", muatanSeharusnyaJikaVolumeBoxDanVolumeGudang0, " ///")
	}
}

// 7. Muatan Gudang - volumeBox > volumeGudang

var (
	muatanSeharusnyaJikaVolumeBoxDiatasVolumeGudang float64 = 0
)

func TestHitungMuatanVolumeBoxDiatasVolumeGudang(t *testing.T) {

	sisiBox = 10
	sisiGudang = 5
	muatan = HitungMuatanGudang(sisiBox, sisiGudang)

	if muatan == muatanSeharusnyaJikaVolumeBoxDiatasVolumeGudang {
		t.Log("/// Muatan : ", muatan, " Expected: ", muatanSeharusnyaJikaVolumeBoxDiatasVolumeGudang, " ///")
	} else {
		t.Error("/// SALAH! Muatan: ", muatan, " Expected: ", muatanSeharusnyaJikaVolumeBoxDiatasVolumeGudang, " ///")
	}

}

// 8. Muatan Gudang - Sukses

var (
	muatanSeharusnyaJikaSukses float64 = 27
)

func TestHitungMuatanSukses(t *testing.T) {

	sisiBox = 3
	sisiGudang = 9
	muatan = HitungMuatanGudang(sisiBox, sisiGudang)

	if muatan == muatanSeharusnyaJikaSukses {
		t.Log("/// Muatan : ", muatan, " Expected: ", muatanSeharusnyaJikaSukses, " ///")
	} else {
		t.Error("/// SALAH! Muatan: ", muatan, " Expected: ", muatanSeharusnyaJikaSukses, " ///")
	}

}

// 9. Total Bayar - Total Diatas 500.000.000

var (
	totalBayarSeharusnyaDiatas500Jt float64 = 720000000
)

func TestTotalBayarDiatas500Jt(t *testing.T) {

	sisiGudang = 20
	totalBayar = HitungTotalBayarSewa(sisiGudang)

	if totalBayar == totalBayarSeharusnyaDiatas500Jt {
		t.Log("/// Total Bayar : ", totalBayar, " Expected: ", totalBayarSeharusnyaDiatas500Jt, " ///")
	} else {
		t.Error("/// SALAH! Total Bayar: ", totalBayar, " Expected: ", totalBayarSeharusnyaDiatas500Jt, " ///")
	}

}

// 10. Total Bayar - Total Dibawah / Sama Dengan 500.000.000
var (
	totalBayarSeharusnyaDibawah500Jt float64 = 6400000
)

func TestTotalBayarDibawah500Jt(t *testing.T) {

	sisiGudang = 4
	totalBayar = HitungTotalBayarSewa(sisiGudang)

	if totalBayar == totalBayarSeharusnyaDibawah500Jt {
		t.Log("/// Total Bayar : ", totalBayar, " Expected: ", totalBayarSeharusnyaDibawah500Jt, " ///")
	} else {
		t.Error("/// SALAH! Total Bayar: ", totalBayar, " Expected: ", totalBayarSeharusnyaDibawah500Jt, " ///")
	}
}
