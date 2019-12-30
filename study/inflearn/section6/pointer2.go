//자료형 : 포인터(2)

package main

import "fmt"

func main() {
	//예제1
	i := 7
	p := &i
	fmt.Println("ex1:", i, *p, &i, p)

	*p++ // 1증가
	fmt.Println("ex1:", i, *p, &i, p)

	*p = 77777 //포인터 변수 역 참조값 변경
	fmt.Println("ex1:", i, *p, &i, p)

	i = 77
	fmt.Println("ex1:", i, *p, &i, p)

}
