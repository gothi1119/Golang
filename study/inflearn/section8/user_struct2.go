//사용자 정의 타입(1)

package main

import "fmt"

type cnt int

func main() {
	//기본 자료형 사용자 정의 타입
	//예제1
	a := cnt(15)
	fmt.Println("ex1:", a)

	//예제2
	var b cnt = 15
	fmt.Println("ex2:", b)

	testConverT(int(b)) //cnt type의 경우 형태가 달라 출력이 안됨, 출력을 위해선 강제형변환 필요
	testConverD(b)
}

func testConverT(i int) {
	fmt.Println("ex3:(Default Type) : ", i)
}

func testConverD(i cnt) {
	fmt.Println("ex4:(Custom Type) : ", i)
}
