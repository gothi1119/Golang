package main

import (
	"fmt"
)

type fileHash struct {
	path string
	hash uint64
	size int64
	err  error
}
func main() {
	fmt.Println([][]fileHash)
}
