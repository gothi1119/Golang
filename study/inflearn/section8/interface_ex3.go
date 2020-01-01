//인터페이스 고급(3)

package main

import "fmt"

func checkType(arg interface{}) {
	//arg.(type)을 통해 현재 데이터형 반환
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool :", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int :", arg)
	case float64:
		fmt.Println("This is a float", arg)
	case string:
		fmt.Println("This is a string", arg)
	case nil:
		fmt.Println("This is a nil", arg)
	default:
		fmt.Println("What type is it?", arg)
	}
}

func main() {
	//실제 타입 검사 switch
	//빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크를 통해 형 변환 후 사용 가능

	checkType(true)
	checkType(1)
	checkType(22.542)
	checkType(nil)
	checkType("Hello Golang")
}
