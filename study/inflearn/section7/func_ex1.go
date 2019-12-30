//함수 심화(1)

package main

import "fmt"

func multifly(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}
	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2:", value)
	}
}

func main() {
	//함수 고급
	//가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)
	x := multifly(1, 2, 3, 4, 5)
	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("ex1:", x)
	fmt.Println("ex2:", y)

	prtWord("a", "apple", "b", "seoul", "good")

	fmt.Println()
	//예제3
	a := []int{5, 6, 7, 8, 9, 10}

	m := multifly(a...)
	n := sum(a...)

	fmt.Println("ex3:", m)
	fmt.Println("ex3:", n)

}
