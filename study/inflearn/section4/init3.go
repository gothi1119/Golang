//go 초기화 함수(3)

package main

import (
	"fmt"
	"section4/lib"
)

var num int32

func init() {
	num = 30
	fmt.Println("Main init start!")
}
func main() {
	fmt.Println("10보다 큰 수? : ", lib.CheckNum(num))
}
