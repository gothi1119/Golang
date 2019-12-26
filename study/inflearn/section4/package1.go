//패키지(1)

package main

//선언방법
/*
import "fmt"
import "os"
*/
import (
  "fmt"
  "os"
)

func main(){
  //패키지 : 코드 모듈화 및 재사용
  //응집도, 결합도
  //Go : 패키지 단위의 독립이고 작은 단위로 개발 -> 작은 패키지를 결합하여 프로그램 작성 권고
  // 패키지 이름 = 디렉토리 이름
  // 같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
  // 네이밍 규칙 : 소문자 Private, 대문자 : Public

  var name string

  fmt.Println("이름은? : ")
  fmt.Scanf("%s",&name)

  fmt.Fprintf(os.Stdout, "Hi %s\n",name)
}
