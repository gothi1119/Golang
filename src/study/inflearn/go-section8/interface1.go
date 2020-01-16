//인터페이스 기본(1)
package main

import "fmt"

type test interface{} //빈 인터페이스

func main() {
	//인터페이스
	//객체의 동작을 표현, 골격
	//단순히 동작에 대한 방법만 표시
	//추상화 제공
	//인터페이스의 메소드를 구현한 타입은 인터페이스로 사용 가능
	//Go언어를 유연하게 사용 가능
	//덕타이핑 : Go 언어의 독창적인 특성

	/*
	  type 인터페이스명 interface {
	     메서드1() 반환 값(타입형)
	  	 메서드2() //반환 값 없을 경우
	   }
	*/

	//예제1
	var t test
	fmt.Println("ex1 : ", t) //Empty 인터페이스인 경우 nil 리턴
}
