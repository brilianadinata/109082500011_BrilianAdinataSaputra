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