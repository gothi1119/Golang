//구조체 심화(2)

package main

import "fmt"

type account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	//예제1
	kim := account{"245-901", 10000000, 0.015}
	lee := account{"245-901", 14000000, 0.035}

	fmt.Println("ex1: ", kim)
	fmt.Println("ex1: ", lee)
	fmt.Println()

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println("ex2: ", int(kim.balance))
	fmt.Println("ex2: ", int(lee.balance))

}
