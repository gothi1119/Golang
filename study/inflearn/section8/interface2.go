////인터페이스 기본(2)

package main

import "fmt"

type dog struct {
	name   string
	weight int
}

//bite Method
func (d dog) bite() {
	fmt.Println(d.name, " bites!")
}

//동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {
	//인터페이스 구현 예제
	//예제1
	dog1 := dog{"paul", 10}
	var interface1 Behavior
	interface1 = dog1
	interface1.bite()
	//dog1.bite()

	//예제2
	dog2 := dog{"marry", 12}
	inter2 := Behavior(dog2) //이게 더 많이 쓰임
	inter2.bite()

	//예제3
	inters := []Behavior{dog1, dog2}

	//인덱스 형태로 실행
	for idx, _ := range inters {
		inters[idx].bite()
	}

	//값 형태로 실행 (인터페이스)
	for _, val := range inters {
		val.bite()
	}
}
