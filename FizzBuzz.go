package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := make([]string, 0, 30)
	for i := 1; i <= 30; i++ {
		result = append(result, fizzbuzz(i))
	}
	fmt.Println(result)
}

func fizzbuzz(x int) string{
	if x%3 == 0 && x%5 == 0 {
		return "FizzBuzz"
	} else if x%3 == 0 {
		return "Fizz"
	} else if x%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(x)
	}
}
