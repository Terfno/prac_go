package main

import "fmt"

func main() {
	// add
	fmt.Println("3 + 4 =", adder(3, 4))

	// swap
	a, b := swap("hello", "world")
	fmt.Println("swap! 'hello' and 'world':", a, b)

	// naked return
	c, d := kakewari(4, 2)
	fmt.Println("掛け算:", c, "\n割り算:", d)
}

func adder(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func kakewari(x, y int) (a, b int) {
	a = x * y
	b = x / y
	return
}
