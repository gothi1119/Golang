package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

func main() {

	dirname := "C:\\Program Files\\AhnLab" + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Reading " + dirname)

	for _, file := range files {
		fmt.Println(reflect.TypeOf(file))
		/*if file.Mode() {
		if filepath.Ext(file.Name()) == "." {
			fmt.Println("file.Name()")
			//fmt.Println("Deleted ", file.Name())
		*/
	}
}
