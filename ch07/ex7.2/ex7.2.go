package main

import "fmt"

func main() {
	math := 80
	eng := 74
	history := 95
	fmt.Println("김일등 님 평균 점수는", (math+eng+history)/3, "입니다.")

	math = 88
	eng = 92
	history = 53
	fmt.Println("김이등 님 평균 점수는", (math+eng+history)/3, "입니다.")
}
