package main

import (
	"fmt"
	"math"
)

const i = 9

func compare(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(compare(2, 10, 20))
}
