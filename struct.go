package main

import "fmt"

type Intro struct { // 自己紹介のための構造体
	age   int
	name  string
	blood string
}

func main() {
	var (
		sue = Intro{18, "Takahito Sueda", "O"}
		miya = Intro{age: 18, name: "Hikaru Miyatani"}
		anyone = Intro{name: "ANYONE"}
	)

	fmt.Println(sue, miya, anyone)
}
