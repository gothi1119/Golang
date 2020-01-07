package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ph := "C:/Users/HM-Fornesic2/Desktop"
	//var files []string
	err := filepath.Walk(ph, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.Size())
		return nil
	})
	if err != nil {
		panic(err)
	}
	/*for _, file := range files {
		fmt.Println(file)
	}*/
}
