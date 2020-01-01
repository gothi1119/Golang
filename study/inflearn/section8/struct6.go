//구제초 기본(5)

package main

import "fmt"

type Car struct { //첫글자 대문자로 선언
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct { //첫글자 소문자로 선언
	length int "전장"
	heigth int "전고"
	width  int "전축"
}

func main() {
	//중첩 구조체
	//예제1

	car1 := Car{
		"520d",
		"Silver",
		"bmw",
		spec{4000, 1000, 2000},
	}

	// 내부 spec 구조체 출력
	fmt.Println("ex1 :", car1.name)
	fmt.Println("ex1 :", car1.color)
	fmt.Println("ex1 :", car1.company)
	fmt.Printf("ex1: %#v\n", car1.detail)

	//예제2
	//내부 spect 구조체 필드 값
	fmt.Println("ex2:", car1.detail.length)
	fmt.Println("ex2:", car1.detail.heigth)
	fmt.Println("ex2:", car1.detail.width)

}
