package main

import "fmt"

func main() {
	a := [3]int{1,2,3}

	fmt.Println(a)

	a[1] = 5
	fmt.Println(a)
}
