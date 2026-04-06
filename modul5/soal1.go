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