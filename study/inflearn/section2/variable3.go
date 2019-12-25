//변수3
package main

import "fmt"

func main() {
	//짧은 선언
	//반드시 한수 안에서만 사용(전역변수가 아님, 지역변수), 선언 후 재할당시 예외 발생
	//주로 제한된 범위의 함수내에서 사용할 경우 코드 가독성을 높일 수 있음

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	//shortVar3 := true 예외 발생

	fmt.Println("shortVar1: ", shortVar1, "shortVar2: ", shortVar2, "shortVar3: ", shortVar3)

	//Examples
	if i := 10; i < 11 {
		fmt.Println("ShotVariable Success")

	}
}
