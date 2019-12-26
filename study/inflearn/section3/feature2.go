//GO 특징(2)
package main

import "fmt"

func main(){
  //문장 끝 세미콜론(;) 주의
  //자동으로 끝에 세미콜론 삽입
  //두 문장을 한 문장에 표현할 경우 명시적으로 세미적으로 사용 필요 (비권장)
  //반복문 및 제어문(조건문)(if,for) 사용할 경우에 주의

  //예제1
  for i :=0; i<=10; i++ {
    fmt.Println("ex1:",i)
    fmt.Println(i)
    }

  //예제2
  for j:=10; j>0; j--{
    fmt.Print("ex2:",j)
  }
}
