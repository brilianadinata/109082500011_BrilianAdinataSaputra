# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Brilian Adinata Saputra - 109082500011</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
	satu, dua, tiga string
	temp string
)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul2/output/output-soal1.png)
Program diatas adalah implementasi dari code yang sudah diberikan untuk dianalisis, jadi untuk program diatas saya disuruh untuk menginput 3 string secara bertahap per kata, lalu ketika sudah selesai menginput 3 kata yang harus diisi maka program akan menampilkan output dari gabungan input an 3 kata yang tadi saya input, dan ada variable "temp" untuk merubah atau membolak balik kata yang sudah saya input. Contoh : Saya inputkan "aku", "orang", "jawa", output akan otomatis menampikan kalimat yang sudah kita input di "ouput awal" dan akan dibalik kata yang sudah input di "output akhir" karena variable dari temp.


### 2. [Soal]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul2/output/output-soal2.png)
Program diatas adalah implementasi dari program yang dibuat untuk melakukan percobaan kimia sebanyak 5 kali percobaan. Pada setiap percobaan kita diminta untuk menginput 4 warna cairan yang ada di tabung reaksi secara berurutan. Program kemudian akan mengecek apakah warna yang kita input sudah sesuai dengan urutan yang ditentukan yaitu merah, kuning, hijau, dan ungu. Jika dari percobaan pertama sampai percobaan kelima urutan warna yang dimasukkan selalu benar, maka program akan menampilkan hasil true. Tetapi jika ada salah satu percobaan yang warnanya tidak sesuai dengan urutan tersebut, maka hasil yang ditampilkan akan menjadi false.



### 3. [Soal]
#### soal3.go

```go
package main
package main

import "fmt"

func main() {
	var berat int
	var kg int
	var gram int
	var biayaKG int
	var biayaGR int
	var total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	kg = berat / 1000
	gram = berat % 1000

	biayaKG = kg * 10000

	if gram < 500 {
		biayaGR = gram * 5
	} else {
		biayaGR = gram * 15
	}

	total = biayaKG + biayaGR

	if berat > 10000 {
		total = biayaKG
	}

	fmt.Println("Detail berat:", kg, "kg +", gram, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKG, "+ Rp.", biayaGR)
	fmt.Println("Total biaya: Rp.", total)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul2/output/output-soal3.png)
Program diatas adalah implementasi dari program yang dibuat untuk menghitung biaya pengiriman parsel berdasarkan berat yang kita input dalam satuan gram. Setelah berat dimasukkan, program akan membagi berat tersebut menjadi kilogram dan sisa gram. Setiap 1 kilogram dikenakan biaya Rp 10.000, lalu jika masih ada sisa gram maka akan dihitung biaya tambahan. Jika sisa gram kurang dari 500 gram maka biayanya Rp 5 per gram, sedangkan jika lebih dari atau sama dengan 500 gram maka biayanya Rp 15 per gram. Tetapi jika berat parsel lebih dari 10 kg, maka biaya dari sisa gram tidak dihitung dan hanya biaya kilogram saja yang dihitung sebagai total biaya pengiriman.
