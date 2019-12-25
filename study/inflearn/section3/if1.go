//if문 1
package main

import "fmt"

func main() {
	//제어문(조건문)
	// IF 문 : 반드시 Boolean으로 검사 -> 1,0 (사용불가 : 자동 형 변환 불가)
	// 소괄호 미사용
	var a int = 20
	b := 20

	//예제1
	if a >= 15 {
		fmt.Println("15이상")
	}
	if b >= 25 {
		fmt.Println("25이상")
	}

	// 에러발생 1
	/*
	  if b>=25
	  {

	  }
	*/

	// 에러발생 2
	/*
	  if b >= 25
	    fmt.Println("Test")
	*/

	// 에러발생 3
	/*
	  if c:=1; c{
	    fmt.Println("Test")
	  }
	*/
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}
	// c+= 20 선언시 에러 발생
}
