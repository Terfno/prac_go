// １から３０までの数字を改行付きで出力してください。 ただし、

// ３の倍数の場合は、上記の代わりに"foo"という文字列を出力
// ５の倍数の場合は、上記の代わりに"bar"という文字列を出力
// ３の倍数かつ５の倍数の場合は、上記の代わりに"foobar"という文字列を出力
package main

import "fmt"

func main() {
	for i := 1; i <= 30; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(x int) {
	if x%3 == 0 && x%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if x%3 == 0 {
		fmt.Println("Fizz")
	} else if x%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(x)
	}
}

// 1
// 2
// Fizz
// 4
// Buzz
// Fizz
// 7
// 8
// Fizz
// Buzz
// 11
// Fizz
// 13
// 14
// FizzBuzz
// 16
// 17
// Fizz
// 19
// Buzz
// Fizz
// 22
// 23
// Fizz
// Buzz
// 26
// Fizz
// 28
// 29
// FizzBuzz