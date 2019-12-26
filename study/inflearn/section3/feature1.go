//Go의 특징(1)

package main

import "fmt"

func main(){
  // Go : 모호하거나, 애매한 문법 금지
  // 후치 연산자 허용 (i++), 전치 연산자 비허용 ++i --> 문법 에러
  // 증감연산 반환값이 없음 ex) sum := i++
  //포인터(Pointer 허용, 연산 비허용)
  // 주석

  //예제1
  sum, i := 0,0

  for i <= 100{
    //sum += i++ 에러
    sum += i
    i++
    // ++i 전위 증감 에러 발생
  }
  fmt.Println("ex1:",sum)
}
