# <h1 align="center">Laporan Praktikum Modul 5 - Rekursif</h1>
<p align="center">Brilian Adinata Saputra - 109082500011</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int

	fmt.Print("Masukkan suku ke-n: ")
	fmt.Scan(&n)

	fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul5/output/output-soal1.png)
Program ini dibuat untuk mengimplementasikan fungsi rekursif dalam menghitung deret Fibonacci. User diminta memasukkan sebuah bilangan bulat yang menyatakan suku ke-n. Fungsi rekursif akan bekerja dengan membagi masalah menjadi sub-masalah yang lebih kecil sesuai rumus Sn = Sn−1 + Sn−2 hingga mencapai base case (0 atau 1). Hasil akhirnya adalah tampilan deret angka mulai dari suku ke-0 hingga suku yang ditentukan oleh user.


### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	cetakBintang(n - 1)
}

func buatPola(n int) {
	if n == 0 {
		return
	}
	buatPola(n - 1)
	
	cetakBintang(n)
	fmt.Println()
}

func main() {
	var input int
	fmt.Scan(&input)
	
	buatPola(input)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul5/output/output-soal2.png)
Program ini dibuat untuk menampilkan pola bintang segitiga siku-siku berdasarkan input N dari user dengan menggunakan metode rekursif. Terdapat dua fungsi rekursif yang digunakan: satu untuk mengatur jumlah baris dan satu lagi untuk mencetak jumlah bintang di setiap barisnya. Dengan menempatkan pemanggilan rekursif sebelum perintah cetak, program akan menghasilkan pola yang meningkat (dari 1 bintang hingga N bintang) sesuai dengan contoh masukan dan keluaran yang diberikan.

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

func cariFaktor(n int, i int) {

	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Printf("%d ", i)
	}

	cariFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	cariFaktor(n, 1)
	fmt.Println()
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul5/output/output-soal3.png)
Program ini dibuat untuk menampilkan seluruh faktor dari sebuah bilangan bulat positif N secara terurut menggunakan fungsi rekursif. Fungsi akan melakukan pengecekan mulai dari angka 1 hingga N dengan memanfaatkan operator modulo untuk menentukan apakah suatu angka merupakan faktor atau bukan. Setiap kali ditemukan angka yang habis membagi N, angka tersebut akan langsung ditampilkan sebelum fungsi memanggil dirinya sendiri kembali dengan nilai pembagi yang bertambah satu.

