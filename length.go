package main

import (
	"fmt"
	"reflect"
)

func main(){
	array := [8]int{1,2,3,4,5,6,7,8}
	length := len(array)
	fmt.Println(array, length, reflect.TypeOf(length))

	s := array[2:4]
	length = len(s)
	capacity := cap(s)
	fmt.Println(s, length, capacity)

	s2 := array[4:6]
	length = len(s2)
	capacity = cap(s2)
	fmt.Println(s2, length, capacity)
}
