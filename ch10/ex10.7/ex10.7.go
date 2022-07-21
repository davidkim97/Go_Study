package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	// age 변수 선언 및 초기화
	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is ", age) // age 값 사용
	}
}
