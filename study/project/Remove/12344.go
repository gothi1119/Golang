package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	dirname := "C:\\Users\\HM-Fornesic2\\Desktop\\Test1234" + string(filepath.Separator)

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
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".txt" {
				err := os.Remove(file.Name())
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Deleted ", file.Name())
			}
		}
	}
}
