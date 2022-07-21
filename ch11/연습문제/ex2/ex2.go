// 주석 부분을 채워서 3단부터 6단까지 출력하지 않도록
package main

import "fmt"

func main() {
	for i := 2; i < 10; i++ {
		// 여기를 채워보시오
		if i >= 3 && i <= 6 {
			continue
		}
		for j := 1; j < 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println()
	}
}
