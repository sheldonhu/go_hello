package main

import "fmt"

func add(x, y int) int {
	c := x + y
	return c
}

func main() {
	fmt.Println(add(1, 2))
}
