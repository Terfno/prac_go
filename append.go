package main

import "fmt"

func main() {
	slice := make([]int, 0, 3)
	showS("元のスライス", slice)

	appendedSlice := append(slice, 1)
	showS("1を追加したスライス", appendedSlice)

	appendedSlice = append(appendedSlice, 2, 3)
	showS("2と3を追加したスライス", appendedSlice)
}

func showS(s string, x []int) {
	fmt.Println(s, "| length=", len(x), "| capacity=", cap(x), "| スライス=", x)
}
