//구조체 기본(4)

package main

import "fmt"

func main() {
	//구조체 익명 선언 및 활용
	//예제1 type 구조체명 타입
	car1 := struct{ name, color string }{"520d", "white"}

	fmt.Println("ex1:", car1)
	fmt.Printf("ex1: %#v\n", car1)

	//예제2
	cars := []struct{ name, color string }{{"520d", "white"}, {"530i", "red"}, {"528i", "blue"}}

	for _, c := range cars {
		fmt.Printf("ex2:(%s, %s)----------%#v\n", c.name, c.color, c)
	}
}
