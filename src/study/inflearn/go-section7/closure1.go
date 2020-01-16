//함수 Closure(1)
package main

import "fmt"

func main() {
	//클로저(Closure)
	//익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	//이 때, 함수는 일급 객체 이므로 변수의 값으로 사용 가능
	//현재 범위에 있는 변수의 값을 캡처 후 호출 할 때 변수 사용 가능(선언 시점 기준)

	//예제1
	multiply := func(x, y int) int { //익명함수 변수 할당
		return x * y
	}

	r1 := multiply(5, 10)

	fmt.Println("ex1 : ", r1)

	//예제2
	m, n := 10, 5            //변수를 캡처
	sum := func(c int) int { //익명함수 변수 할당
		return m + n + c //지역 변수 소멸되지 않는다.(함수 호출 시 마다 사용 가능)
	}

	r2 := sum(10)

	fmt.Println("ex2 : ", r2)

}
