package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	berhasil := true

	for i := 1; i <= 5; i++ {

		fmt.Println("Percobaan", i)

		fmt.Print("Warna 1: ")
		fmt.Scanln(&warna1)

		fmt.Print("Warna 2: ")
		fmt.Scanln(&warna2)

		fmt.Print("Warna 3: ")
		fmt.Scanln(&warna3)

		fmt.Print("Warna 4: ")
		fmt.Scanln(&warna4)

		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}