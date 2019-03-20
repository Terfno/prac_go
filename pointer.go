package main

import "fmt"

func main() {
	i := 1
	j := 2
	fmt.Println(i, j)

	pI := &i // pointer for `i`
	pJ := &j // pointer for `j`

	*pI = 3
	*pJ = 5

	fmt.Println(*pI, *pJ)
	fmt.Println(i, j)
}
