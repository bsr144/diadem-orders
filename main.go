package main

import "fmt"

func main() {
	fmt.Println(numberOfWays(2, 2))
	// fmt.Println(factorial(5))
}

func numberOfWays(n, k int) int {
	if n == 0 {
		return 1
	}
	if n < 0 || k <= 0 {
		return 0
	}

	return numberOfWays(n-k, k) + numberOfWays(n, k-1)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}
