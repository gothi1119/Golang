//함수 심화(4)

package main

import "fmt"

func main() {
	//함수 고급
	//익명 함수(Anonymous Function)
	//즉시 실행 (선언과 동시에)

	//예제1
	func() {
		fmt.Println("ex1: Anonymous function!")
	}()

	//예제2
	msg := "Hello Golang"

	func(m string) {
		fmt.Println("ex2: ", m)
	}(msg)

	//예제3
	func(x, y int) {
		fmt.Println("ex3:", x+y)
	}(10, 20)

	//예제4
	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println("ex4:", r)
}
