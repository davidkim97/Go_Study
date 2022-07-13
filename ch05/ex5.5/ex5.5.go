package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b) // 입력 두개 받기
	if err != nil {            // 에러 발생 시 에러 코드 출력
		fmt.Println(n, err)
	} else { // 정상 입력되면 입력 값 출력
		fmt.Println(n, a, b)
	}
}
