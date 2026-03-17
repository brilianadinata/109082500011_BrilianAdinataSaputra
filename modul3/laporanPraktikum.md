# <h1 align="center">Laporan Praktikum Modul 3 - AP</h1>
<p align="center">Brilian Adinata Saputra - 109082500011</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func factorial(n int) int {
    hasil := 1
    for i := 1; i <= n; i++ {
        hasil = hasil * i
    }
    return hasil
}

func permutation(n, r int) int {
    return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
    return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
    var a, b, c, d int

	fmt.Print("Masukkan nilai a, b, c, d: ")
    fmt.Scan(&a, &b, &c, &d)

    fmt.Println(permutation(a, c), combination(a, c))
    fmt.Println(permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul3output/output-soal1.png)
Program ini dibuat untuk menghitung permutasi dan kombinasi dari beberapa angka yang diinput oleh user. Di awal, kita diminta memasukkan 4 bilangan yaitu a, b, c, dan d. Setelah itu program akan menghitung permutasi dan kombinasi dari a terhadap c, lalu b terhadap d dengan bantuan fungsi faktorial. Hasilnya nanti ditampilkan dalam dua baris, baris pertama untuk a dan c, dan baris kedua untuk b dan d.


### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
    return x * x
}

func g(x int) int {
    return x - 2
}

func h(x int) int {
    return x + 1
}

func main() {
    var a, b, c int

	fmt.Println("Masukkan nilai a, b, c: ")
    fmt.Scan(&a, &b, &c)

    fmt.Println(f(g(h(a))))

    fmt.Println(g(h(f(b))))

    fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul3/output/output-soal2.png)
Program ini digunakan untuk menghitung hasil dari beberapa fungsi yang digabung (komposisi fungsi). Kita diminta memasukkan tiga angka yaitu a, b, dan c. Setelah itu program akan menghitung hasil dari kombinasi fungsi seperti f(g(h(a))), g(h(f(b))), dan h(f(g(c))). Hasil dari masing-masing perhitungan tersebut akan ditampilkan satu per satu dalam bentuk tiga baris.


### 3. [Soal]
#### soal3.go

```go
package main

import (
    "fmt"
    "math"
)

func jarak(a, b, c, d float64) float64 {
    return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
    return jarak(cx, cy, x, y) <= r
}

func main() {
    var cx1, cy1, r1 float64
    var cx2, cy2, r2 float64
    var x, y float64
	
	fmt.Print("Input lingkaran 1 (cx cy r): ")
    fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Input lingkaran 2 (cx cy r): ")
    fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Input titik (x y): ")
    fmt.Scan(&x, &y)

    in1 := didalam(cx1, cy1, r1, x, y)
    in2 := didalam(cx2, cy2, r2, x, y)

    if in1 && in2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if in1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if in2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/brilianadinata/109082500011_BrilianAdinataSaputra/blob/main/modul3/output/output-soal3.png)
Program ini dibuat untuk mengecek posisi suatu titik terhadap dua lingkaran. Kita diminta memasukkan data pusat dan jari-jari dari dua lingkaran, lalu memasukkan satu titik. Program akan menghitung apakah titik tersebut berada di dalam atau di luar masing-masing lingkaran. Nanti hasilnya akan ditampilkan apakah titik berada di dalam lingkaran 1, lingkaran 2, di dalam keduanya, atau di luar semua lingkaran.