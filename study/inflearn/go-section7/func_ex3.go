//함수 심화(3)
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2 : (", n, ")", "hi!")
	prtHello(n - 1)
}

func main() {
	//함수 고급
	//재귀 함수(Recursion)

	//예제1
	x := fact(7)
	fmt.Println("ex1 :", x)

	//예제2
	prtHello(10)
}
