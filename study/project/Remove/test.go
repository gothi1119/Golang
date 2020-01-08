package main

import (
	"log"
	"os"
)

func main() {
	var test []string
	test = append(test, "C:\\Users\\HM-Fornesic2\\Desktop\\Test1234\\321 - 복사본 (4).txt")
	test = append(test, "C:\\Users\\HM-Fornesic2\\Desktop\\Test1234\\321 - 복사본 (8).txt")
	err := os.Remove(test[1])
	if err != nil {
		log.Fatal(err)
	}
}
