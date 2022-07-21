package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	switch age := getMyAge(); age { // getMyAge() 결과값 비교
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is ", age) // age 값 사용
	}

	//fmt.Println("age is ", age) // Error - age 변수는 사라짐.
}
