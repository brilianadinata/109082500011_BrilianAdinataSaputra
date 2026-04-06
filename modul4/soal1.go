package main

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