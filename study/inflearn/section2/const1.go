//상수1
package main

import "fmt"

func main() {
	//상수
	//const 사용 초기화, 한 번 선언 후 값 변경 금지, 고정된 값을 이용한 관리용
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	//const d=getHeight() 함수 리턴값이 매번 다르므로 사용 불가능
	const e = 35.6
	const f = false
	/*
	  에러발생
	  const g string
	  g ="Test3"
	*/

	fmt.Println("a:", a, "b:", b, "c:", c, "e:", e, "f:", f)
}
