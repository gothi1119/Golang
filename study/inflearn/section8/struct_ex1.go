//구조체 심화(1)

package main

import "fmt"

type account struct {
	number   string
	balance  float64
	interest float64
}

//생성자 패턴
func NewAccount(number string, balance float64, interest float64) *account { //포인터 반환이 아닌 경우 값 복사
	return &account{number, balance, interest}
}

func main() {
	//구조체 생성자 패턴 예제

	//예제1
	kim := account{number: "245-901", balance: 1000000, interest: 0.015}

	var lee *account = new(account)
	lee.number = "245-902" //getter,setter
	lee.balance = 13000000
	lee.interest = 025

	fmt.Println("ex1:", kim)
	fmt.Println("ex1:", lee)

	//예제
	park := NewAccount("245-903", 17000000, 0.04)
	fmt.Println("ex2:", park)
}
