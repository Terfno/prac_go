package main

import "fmt"

func main() {
	// add
	fmt.Println("3 + 4 =", adder(3, 4))

	// swap
	a, b := swap("hello", "world")
	fmt.Println("swap! 'hello' and 'world':", a, b)
}

func adder(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}
