//사용자 정의 타입(3)
package main

import "fmt"

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt: %d, price: %d, orderPrice: %d", cnt, price, fn(cnt, price))
}

func main() {
	//함수 사용자 정의 타입

	//예제1
	var orderPrice totCost
	orderPrice = func(cnt int, price int) int {
		return (cnt * price) + 1000000
	}
	describe(3, 10000000, orderPrice)
}
