////인터페이스 기본(4)

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

//구조체 dog 메소드 구현

func (d dog) running() {
	fmt.Println(d.name, " : Dog is running")
}

func (c cat) running() {
	fmt.Println(c.name, " : Cat is running")
}

func act(animal interface{ running() }) { //익명선언
	animal.running()
}

func main() {

	//익명 인퍼테이스 사용 예제(즉시 선언 후 사용)

	//예제1
	dog1 := dog{"paul", 10}
	cat1 := cat{"bob", 5}

	//개 행동 실행
	act(dog1)
	//고양이 행동 실행
	act(cat1)
}
