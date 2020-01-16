//사용자 정의 타입(1)
package main

import "fmt"

// 사용자 정의 타입(구조체)
type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

//구조체 <-> 메소드 바인딩
func (c Car) Price() int64 {
	return c.price + c.tax
}

func main() {
	//Go -> 객체 지향 타입을 구조체로 정의한다.(클래스, 상속 개념 없음)
	//상태, 메소드를 분리해서 정의(결합성 없음)
	//사용자 정의 타입 : 구조체 , 인터페이스 , 기본타입 , 함수
	//구조체와 -> 메서드 연결을 통해서 타 언어의 클래스 형식 처럼 사용 가능(객체 지향)

	//예제1
	bmw := Car{name: "520d", price: 54500000, color: "white", tax: 545000}
	benz := Car{name: "220d", price: 74500000, color: "white", tax: 745000}

	bmwPrice := bmw.Price()
	benzPrice := benz.Price()

	fmt.Println("ex1 : ", &bmw)
	fmt.Println("ex1 : ", bmwPrice)
	fmt.Println("ex1 : ", bmw.Price())

	fmt.Println("ex1 : ", &benz == &bmw)

	fmt.Println("ex1 : ", &benz)
	fmt.Println("ex1 : ", benzPrice)
	fmt.Println("ex1 : ", benz.Price())

}
