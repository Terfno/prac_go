package main

import "fmt"

func main() {
	a := 3
	b := 6
	fmt.Println("aは", kigu(a))
	fmt.Println("bは", kigu(b))
}

// 引数が奇数か偶数かを判定する関数。文字列を返す。
func kigu(x int) string {
	if x%2 == 0 {
		return "偶数"
	}
	return "奇数"
}
