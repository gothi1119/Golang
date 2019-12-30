//함수 심화(3)

package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func printhello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2 : (", n, ")", "hi")
	printhello(n - 1)
}

func main() {
	//함수 고급
	//재귀 함수(Recusion)
	//프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이 : 장점
	//코드가 이해하기 어렵고, 기억공간을 많이 사용, 무한루프 가능성 존재

	//예제1
	x := fact(5)
	fmt.Println(x)

	//예제
	printhello(10)
}
