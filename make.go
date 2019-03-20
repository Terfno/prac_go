package main

import "fmt"

func main() {
	sa := make([]int, 5)
	showS("a", sa)

	sb := make([]int, 0, 5)
	showS("b", sb)

	sc := sb[0:2]
	showS("c", sc)

	sd := sb[1:5]
	showS("d",sd)
}

func showS(s string, x []int) {
	fmt.Println(s, "| length=", len(x), "| capacity=", cap(x), "| s=", x)
}
