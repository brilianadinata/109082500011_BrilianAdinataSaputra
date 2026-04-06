package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenangNama string
	var soalAktif, skorAktif int
	var maxSoal, minSkor int

	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&soalAktif, &skorAktif)

		if soalAktif > maxSoal || (soalAktif == maxSoal && skorAktif < minSkor) {
			maxSoal = soalAktif
			minSkor = skorAktif
			pemenangNama = nama
		}
	}

	if pemenangNama != "" {
		fmt.Printf("%s %d %d\n", pemenangNama, maxSoal, minSkor)
	}
}