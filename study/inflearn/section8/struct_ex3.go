//구조체 심화(3)

package main

import "fmt"

type account struct {
	number   string
	balance  float64
	interest float64
}

func (a account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *account) CalculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	//정리 : 구조체 인스턴스 값 변경 시 -> 포인터 전달, 보통의 경우 -> 값 전달
	kim := account{"245-901", 10000000, 0.015}
	lee := account{"245-902", 14000000, 0.015}

	fmt.Println("ex1: ", kim)
	fmt.Println("ex1: ", lee)
	fmt.Println()

	kim.CalculateD(15000000)
	lee.CalculateP(10000000)

	fmt.Println("ex2: ", int(kim.balance))
	fmt.Println("ex2: ", int(lee.balance))
}
