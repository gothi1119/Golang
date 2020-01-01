// 구조체 기본(2)

package main

import "fmt"

type account struct {
	number   string
	balance  float64
	interest float64
}

func (a account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	//다양한 선언 방법
	//&struct, struct  : &struct 포인터를 받아오고 역참조를 또 하기 때문에 속도는 조금 느리다.
	//포인터 구조체를 사용하는 경우 : 인터페이스 메소드를 선언만 해둔 후 -> 오버라이딩 해서 메서드에 포인터 리시버를 사용할 경우에는 반드시 &struct 타입으로 넘겨야함

	//선언방법 1
	var kim *account = new(account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	//선언방법 2
	hong := &account{number: "245-902", balance: 150000000, interest: 0.04}

	lee := new(account)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1:", kim)
	fmt.Println("ex1:", hong)
	fmt.Println("ex1:", lee)

	fmt.Printf("ex2:%#v\n", kim)
	fmt.Printf("ex2:%#v\n", hong)
	fmt.Printf("ex2:%#v\n", lee)
	fmt.Println()

	//예제2
	fmt.Println("ex3:", int(kim.Calculate()))
	fmt.Println("ex3:", int(lee.Calculate()))
	fmt.Println("ex3:", int(hong.Calculate()))

}
