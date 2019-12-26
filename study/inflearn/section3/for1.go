//for문(1)
package main

import "fmt"

func main(){
  //반복문 - for
  //Go에서 유일하게 제공되는 반복문
  //다양한 사용법 숙지

  //예제1

  for i:=0; i<5; i++{
    fmt.Println(i)
  }
  //에러발생1
  /*
  for i:=0; i<5; i++
  {
    fmt.Println("Test")
  }
  */

  //에러발생2
  /*
  for i:=0;i<5;i++
    fmt.Println("Test")
  */

  //예제2(무한 루프)
 /*
  for {
    fmt.Println("ex2 : Hello, Golang")
    fmt.Println("ex2 : Infinite Loop!")
  }
  */

  //예제3 (Range 용법)
  loc := []string{"Seoul","Busan","Incheon"}
  for index, name := range loc{
    fmt.Println("ex3 : ", index,name)
  }

}
