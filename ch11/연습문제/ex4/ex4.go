/*
 이중 for문을 사용해서 다음 별 모양 출력
 *****
 ****
 ***
 **
 *

*/

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
