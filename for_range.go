package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	for i, v := range array{
		array[i] = v*v
	}
	fmt.Println(array)
}
