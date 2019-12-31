//사용자 정의 타입(4)

package main

import "fmt"

type shoppingBasket struct{ cnt, price int }

//결제 함수
func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

//원본 수정 (참조형식)
func (b *shoppingBasket) repurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

//원본 수정X (값 전달 형식)
func (b shoppingBasket) repurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	//리시버(구조체 메소드) 전달(값, 참조) 형식
	//함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정x) -> 맵, 슬라이스 참조 전달
	//리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정이 가능

	//예제1
	bs1 := shoppingBasket{3, 5000}
	fmt.Println("ex1(totPrice)", bs1.purchase())

	//참조 전달(원본 값이 변경)
	bs1.repurchaseP(7, 5000)
	fmt.Println("ex1(totPrice)", bs1.purchase())

	//매개 변수 전달 (원본 값이 변경x)
	bs1.repurchaseD(10, 0)
	fmt.Println("ex1(totPrice)", bs1.purchase())

}
