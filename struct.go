package main

import "fmt"

type Intro struct { // 自己紹介のための構造体
	age   int
	name  string
	blood string
}

func main() {
	me := Intro{18, "Takahito Sueda", "O"}
	miya := Intro{17, "Miyatani Hikaru", "B"}

	fmt.Println(me.blood, miya.blood)

	miya.blood = "C"

	fmt.Println(miya.blood)
}
