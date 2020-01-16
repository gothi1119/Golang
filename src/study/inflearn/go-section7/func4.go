//함수 기초(4)
package main

import "fmt"

func multiply(x int, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return //리턴 변수 지정
}

func multiply2(x int, y int) (int, int) {
	return x * 10, y * 20
}

func main() {
	//다중 값 반환(return values)

	//예제1
	a, b := multiply(10, 5)
	fmt.Println("ex1 : ", a, b)

	//예제2
	c, d := multiply2(10, 5)
	fmt.Println("ex2 : ", c, d)
}
