//열거형3
package main

import "fmt"

func main() {
	// _ 변수명을 이용하여 Skip처리 가능
	const (
		_ = iota + 0.75*2
		DEAUFLT
		SILVER
		GOLD
		PLATINUM
	)
	fmt.Println("D:", DEAUFLT)
	fmt.Println("S:", SILVER)
	fmt.Println("G:", GOLD)
	fmt.Println("P:", PLATINUM)
}
