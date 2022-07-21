/*
다음 출력 결과가 나오도록 9까지의 홀수의 제곱값을 출력하시오.
1 * 1 = 1
3 * 3 = 9
5 * 5 = 25
7 * 7 = 49
9 * 9 = 81
*/

package main

import "fmt"

func main() {
	for i := 1; i < 10; i += 2 {
		fmt.Println(i, "*", i, "=", i*i)
	}
}
