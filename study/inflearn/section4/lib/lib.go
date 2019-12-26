package lib

import "fmt"

func init() {
	fmt.Println("lib Package -> init start!")
}

func CheckNum(c int32) bool {
	return c > 10
}
