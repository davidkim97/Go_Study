package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	}

	// 인덱스 무효화
	// for _, v := range t {
	// 	fmt.Println(v)
	// }
}
