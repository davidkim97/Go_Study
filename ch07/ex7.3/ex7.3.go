package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는 ", avg, "점 입니다.")
}

func main() {
	PrintAvgScore("김일등", 80, 74, 95)
	PrintAvgScore("김이등", 88, 92, 53)
}
