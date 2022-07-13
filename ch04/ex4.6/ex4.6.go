package main

import "fmt"

var g int = 10 // 전역 변수

func main() {
	var m int = 20 // 지역변수
	{
		var s int = 50 // 지역변수
		fmt.Println(m, s, g)
	} // s 지역변수는 사라짐.

	m = s + 20
}
