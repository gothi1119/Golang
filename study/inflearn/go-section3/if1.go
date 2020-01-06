//IF문(1)
package main

import "fmt"

func main() {
	//제어문 - IF
	//반드시 Boolean 식으로 검사 -> 1, 0 (자동 형 변환 없음)
	//소괄호 사용 x

	var a int = 20
	b := 30

	//예제1
	if a >= 15 {
		fmt.Println("15 이상")
	}

	//예제2
	if b >= 25 {
		fmt.Println("25 이상")
	}

	//에러 발생1
	/*
		if b >= 25
		{
			fmt.Println("25 이상")
		}
	*/

	//에러 발생2
	/*
		if b >= 25
			fmt.Println("25 이상")
	*/

	//에러 발생3
	/*
		if c := 1; c {
			fmt.Println("True")
		}
	*/

	//예제3 (간단한 문장 : Optional Statement)
	if c := 40; c >= 35 {
		fmt.Println("35 이상")
	}
	//c += 20 //에러 발생
}
