//함수 기초 (4)

package main

import "fmt"

func multiply(x int, y int) (r1, r2 int) { //(x,y int) 가능
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

func multiply2(x int, y int) (int, int) { //(x,y int) 가능
	return x * 10, y * 20
}

func main() {
	//리턴 값 변수 사용
	a, b := multiply(10, 5)
	fmt.Println("ex1:", a, b)

	c, d := multiply2(10, 5)
	fmt.Println("ex1:", c, d)

}
