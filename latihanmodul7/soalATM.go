package main

import "fmt"

type atm struct {
	norek int
	pin   string
	saldo int
}

func main() {
	var rekening atm
	rekening.norek = 889900
	rekening.pin = "123456"
	rekening.saldo = 500000

	var inputrek int
	var inputPin string
	var tarik int

	fmt.Print("Masukkan nomor rekening: ")
	fmt.Scan(&inputrek)

	if inputrek != rekening.norek {
		fmt.Println("error rekening tidak terdaftar")
		return
	}

	fmt.Print("Masukkan PIN: ")
	fmt.Scan(&inputPin)

	if inputPin != rekening.pin {
		fmt.Println("error pin tidak valid")
		return
	}

	fmt.Print("Masukkan jumlah penarikan: ")
	fmt.Scan(&tarik)

	if tarik > rekening.saldo {
		fmt.Println("error saldo tidak mencukupi")
		return
	}

	rekening.saldo -= tarik
	fmt.Println("Penarikan berhasil")
	fmt.Println("Sisa saldo:", rekening.saldo)
}