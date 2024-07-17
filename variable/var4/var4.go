package main

import "fmt"

func main() {
	var a int16 = 3456
	var b int8 = int8(a) // int16을 int8로 변환
	fmt.Println(a, b)
}
