//未説明
package main

import "fmt"

type manji struct {
	Lat, Long float64
}

var m map[string]manji

func main() {
	m = make(map[string]manji)
	m["Bell Labs"] = manji{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"].Lat)
}
