package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

func main() {
	dirname := "." + string(filepath.Separator)
	fmt.Println(dirname)
	d, err := os.Open(dirname)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer d.Close()
	files, err := d.Readdir(2)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	fmt.Println("Reading" + dirname)

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".go" {
				fmt.Println(reflect.TypeOf(files))
			}
		}
	}

}
