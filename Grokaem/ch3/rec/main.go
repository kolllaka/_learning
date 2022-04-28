package main

import "fmt"

func main() {
	// count(3)
	// fmt.Println()

	fmt.Printf("factorial(3): %v\n", factorial(3))
	fmt.Printf("factorial(5): %v\n", factorial(5))
}

func count(n int) {
	fmt.Printf("%d...", n)

	if n-1 > 0 {
		count(n - 1)
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return factorial(n-1) * n
}
