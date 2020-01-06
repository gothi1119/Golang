//함수 Defer(2)
package main

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi ")
	}()
}

func main() {

	//예제1
	sayHello("Golang!")
}
