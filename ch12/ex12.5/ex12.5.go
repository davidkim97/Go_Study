package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a { // 배열 a 요소 출력
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()         // 개행
	for i, v := range b { // 배열 b 요소 출력
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a // 배열 a를 b변수에 복사

	fmt.Println()         // 개행
	for i, v := range b { // 배열 b 요소 출력
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
