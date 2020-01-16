//사용자 정의 타입(2)
package main

import "fmt"

type cnt int

func main() {
	//기본 자료형 사용자 정의 타입

	//예제1
	a := cnt(5)

	fmt.Println("ex1 : ", a)

	//예제2
	var b cnt = 15

	fmt.Println("ex2 : ", b)
	//testConvertT(b) //예외 발생 (중요!) 사용자 정의 타입 <-> 기본 타입 : 매개변수 전달 시에 변환해야 사용 가능 (cnt(5), int(5))
	testConvertT(int(b))

	testConvertD(b)
	testConvertD(cnt(10)) //사용 가능
	//testConvertD(int(b)) //예외 발생 (중요!) 사용자 정의 타입 <-> 기본 타입 : 매개변수 전달 시에 변환해야 사용 가능 (cnt(5), int(5))
}

func testConvertT(i int) {
	fmt.Println("ex 2 : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("ex 2 : (Custom Type) : ", i)
}
