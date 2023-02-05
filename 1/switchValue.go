package main

import "fmt"

func main() {
	a := 50
	b := 63
	fmt.Println("a sebelum = ", a)
	fmt.Println("b sebelum = ", b)

	a, b = b, a
	fmt.Println("a sesudah = ", a)
	fmt.Println("b sesudah = ", b)

}
