//사용자 정의 타입(3)

package main

import "fmt"

type totCost func(int, int) int

func describe(cnt, price int, fn totCost) {
	fmt.Printf("cnt : %d, price : %d, orderprice : %d", cnt, price, fn(cnt, price))
}

func main() {
	//예제1
	var orderPrice totCost
	orderPrice = func(cnt, price int) int {
		return (cnt * price) + 100000
	}
	describe(3, 5000, orderPrice)
}
