# <h1 align="center">Laporan Praktikum Modul 3 - AP</h1>
<p align="center">Brilian Adinata Saputra - 109082500011</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n int, r int, hasil *int) {
	var nFact, nrFact int
	factorial(n, &nFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / nrFact
}

func combination(n int, r int, hasil *int) {
	var pResult, rFact int
	permutation(n, r, &pResult) 
	factorial(r, &rFact)
	*hasil = pResult / rFact
}

func main() {
	var a, b, c, d int
	var resP1, resC1, resP2, resC2 int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &resP1)
	combination(a, c, &resC1)
	fmt.Printf("%d %d\n", resP1, resC1)

	permutation(b, d, &resP2)
	combination(b, d, &resC2)
	fmt.Printf("%d %d\n", resP2, resC2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul4/output/output-soal1.png)
Program ini dibuat untuk menghitung permutasi dan kombinasi dari beberapa angka yang diinput oleh user. Di awal, kita diminta memasukkan 4 bilangan yaitu a, b, c, dan d. Setelah itu program akan menghitung permutasi dan kombinasi dari a terhadap c, lalu b terhadap d dengan bantuan fungsi faktorial. Hasilnya nanti ditampilkan dalam dua baris, baris pertama untuk a dan c, dan baris kedua untuk b dan d.


### 2. [Soal]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul4/output/output-soal2.png)
Program ini digunakan untuk menghitung hasil dari beberapa fungsi yang digabung (komposisi fungsi). Kita diminta memasukkan tiga angka yaitu a, b, dan c. Setelah itu program akan menghitung hasil dari kombinasi fungsi seperti f(g(h(a))), g(h(f(b))), dan h(f(g(c))). Hasil dari masing-masing perhitungan tersebut akan ditampilkan satu per satu dalam bentuk tiga baris.

