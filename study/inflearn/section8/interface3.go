////인터페이스 기본(3)

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

func (d dog) bite() {
	fmt.Println(d.name, " : Dog bites")
}

func (d dog) sounds() {
	fmt.Println(d.name, ": Dog barks")
}

func (d dog) running() {
	fmt.Println(d.name, " : Dog is running")
}

func (c cat) bite() {
	fmt.Println(c.name, " : Cat bites")
}

func (c cat) sounds() {
	fmt.Println(c.name, ": Cat barks")
}

func (c cat) running() {
	fmt.Println(c.name, " : Cat is running")
}

//동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
	sounds()
	running()
}

//인터페이스의 파라미터를 받는다
func act(animal Behavior) {
	animal.bite()
	animal.running()
	animal.sounds()
}

func main() {

	//인터페이스 구현 예제
	//인터페이스 규격화 역할 이해
	//인터페이스에 정의된 메소드 사용 유도
	//코드의 가독성 및 유지보수 증가

	//덕타이핑 예제
	//덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
	//Go의 중요한 특징 : 오리처럼 걷고, 소리내고, 헤엄 등 행동이 같으면 오리라고 볼 수 있다.

	dog1 := dog{"paul", 10}
	cat1 := cat{"bob", 5}

	//개 행동 실행
	act(dog1)
	//고양이 행동 실행
	act(cat1)
}
