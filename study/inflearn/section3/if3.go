//if문(3)
package main

import "fmt"

func main() {
	i := 93

	//if - else if 예제(1)
	if i > 120 {
		fmt.Println("120이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100이상 120미만")
	} else if i < 100 && i >= 50 {
		fmt.Println("100미만 50 이상")
	} else {
		fmt.Println("50이하")
	}

}
