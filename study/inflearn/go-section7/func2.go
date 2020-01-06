//함수 기초(2)
package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10) // add(1, 2)를 호출
}

func add(a, b int) {
	fmt.Println("ex1 :", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i *= 10 // *i = *i * 10
}

func main() {
	//함수(콜백) 및 참조 전달(call by reference) 및 값 전달(call by value)

	//예제1 (콜백 호출)
	sum(10, add) //함수 전달

	//예제2
	//call by value
	a := 100

	multi_value(a)
	fmt.Println("ex2 :", a)

	//예제3
	//reference by value
	b := 100

	multi_reference(&b)
	fmt.Println("ex3 :", b)

}
