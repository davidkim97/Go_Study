package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b) // float64에서 int로 타입 변환
	d := float64(a * c)

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3) // float64 에서 int로 변환
	var h int = int(b) * 3 // float64에서 int로 변환 g 와 값이 다름.

	fmt.Println(g, h, f)

}
