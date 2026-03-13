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