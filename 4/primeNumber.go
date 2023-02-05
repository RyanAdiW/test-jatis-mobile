package main

import "fmt"

func main() {
	n := 100
	count := 0
	for i := 1; i <= n; i++ {
		isPrime := primeNumber(i)
		if i%9 == 0 {
			count++
			fmt.Println("Kelipatan 9 ke-", count)
		} else if isPrime {
			fmt.Println("Bilangan prima: ", i)
		} else {
			fmt.Println(i)
		}
	}
}

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
