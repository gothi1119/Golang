////인터페이스 기본(5)

package main

import "fmt"

type dog struct {
	name   string
	weight int
}

type cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex1:", s)
}

func main() {
	//인터페이스 활용 (빈 인터페이스)
	//함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다.(만능) -> 모든 타입 지정 가능

	dog1 := dog{"paul", 10}
	cat1 := cat{"cat", 5}

	printValue(dog1)
	printValue(cat1)

}
