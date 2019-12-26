//패키지 접근제어(2)
package main

import (
	"fmt"
	checkUp "section4/lib"
	_ "section4/lib2"
)

func main() {
	//패키지 접근제어
	//별칭 사용
	//빈 식별자 사용

	fmt.Println("10보다 큰 수?: ", checkUp.CheckNum(11))
}
