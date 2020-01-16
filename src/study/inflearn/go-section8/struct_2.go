//구조체 기본(2)
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {

	//예제1
	//선언 방법1
	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	//선언 방법2
	hong := &Account{number: "245-902", balance: 15000000, interest: 0.04}

	//선언 방법3
	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", hong)
	fmt.Println("ex1 : ", lee)
	fmt.Printf("ex1 : %#v\n", kim)
	fmt.Printf("ex1 : %#v\n", hong)
	fmt.Printf("ex1 : %#v\n", lee)

	fmt.Println()

	//예제2
	fmt.Println("ex2 : ", int(kim.Calculate()))
	fmt.Println("ex2 : ", int(hong.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))

}
