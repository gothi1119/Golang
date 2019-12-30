//자료형 : 맵(4)

package main

import "fmt"

func main() {
	//맵(Map)
	//맵 조회할 경우 주의 할 점

	//예제1
	map1 := map[string]int{ //int : 0, string : "", float =0.0
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]

	fmt.Println("ex1:", map1)
	fmt.Println("ex1:", value1)
	fmt.Println("ex1:", value2)
	fmt.Println("ex1:", value3, ok) //두번째 리턴 값으로 키 존재 유무 확인

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex:", value)
	} else {
		fmt.Println("kiwi is not exist!")
	}
	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("kiwi is not exist!")
	}
}
