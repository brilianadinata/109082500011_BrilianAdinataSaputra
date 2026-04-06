package main

import "fmt"

type rental struct {
	namapenyewa string
	jenis_naguy string
	tarif_naguy int
}

type waktu struct {
	jam int
	menit int
	detik int
}

func main() {
	var sewa rental
	var waktu waktu

	fmt.Print("Masukkan nama penyewa: ")
	fmt.Scan(&sewa.namapenyewa)

	fmt.Print("Masukkan durasi sewa: ")
	fmt.Scan(&waktu.jam, &waktu.menit, &waktu.detik)

	fmt.Print("Masukkan jenis naguy: ")
	fmt.Scan(&sewa.jenis_naguy)

	if sewa.jenis_naguy == "naga_gembul" {
		sewa.tarif_naguy = 500
	} else if sewa.jenis_naguy == "naga_cungkring" {
		sewa.tarif_naguy = 200
	}
	denda_seblak_ceker == "durasi" {
	}
}