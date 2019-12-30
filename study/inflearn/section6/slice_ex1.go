//자료형 : 슬라이스 심화(1)

package main

import "fmt"

func main() {
	//슬라이스 추가 및 병합
	//예제1

	s1 := []int{1, 2, 3, 4, 5} //7,10  --> 용량이 넘칠경우 새로운 임시 배열(용량 2배)를 만들어 재할당 --> 비효율, 엄청난 스트레스 증가
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      //슬라이스를 삽입할 경우 ...사용
	s3 = append(s2, s3[0:3]...) // 추출 후 병합

	fmt.Println("ex1:", s1)
	fmt.Println("ex1:", s2)
	fmt.Println("ex1:", s3)

	//예제2
	s4 := make([]int, 0, 5)

	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Println("ex2 -> len : %d, cap : %d, value %v \n", len(s4), cap(s4), s4) //길이 및 자동 증가
	}
}
