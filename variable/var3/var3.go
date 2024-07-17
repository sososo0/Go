package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5
	// var c int = b
	// d := a * b

	var c int = int(b)
	d := float64(a) * b

	fmt.Println(c, d)
}
