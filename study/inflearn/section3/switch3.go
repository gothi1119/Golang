//Switch문(3)
package main

import "fmt"

func main() {
	a := 30 / 15
	switch a {
	case 2, 4, 6: //i가 2,4,6 일 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: //i가 1,3,5 일 경우
		fmt.Println("a -> ", a, "는 홀수")
	}

	//예제2
	switch e := "go"; e {
	case "java":
		fmt.Println("Java")
	case "go":
		fmt.Println("Go")
		fallthrough
	case "python":
		fmt.Println("python")
	case "ruby":
		fmt.Println("ruby")
	case "c":
		fmt.Println("C")
	}

}
