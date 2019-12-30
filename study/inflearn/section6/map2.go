//자료형 : 맵(1)

package main

import "fmt"

func main() {
	//맵(Map)
	//맵 조회 및 순회

	//예제1

	map1 := map[string]string{
		"daum":   "https://daum.net",
		"naver":  "https://naver.com",
		"google": "https://googole.com",
	}
	fmt.Println("ex1:", map1["google"])
	fmt.Println("ex1:", map1["daum"])
	fmt.Println()

	//예제2(순서가 없으므로 랜덤)
	for _, v := range map1 {
		fmt.Println("ex2", v)
	}

}
