package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//대상 디렉터리 경로를 입력받음
func input_path(dir string) string {
	fmt.Println("경로 입력:")
	fmt.Scan(&dir)
	return dir
}

//대상 디렉터리 하위 포함한 정보 출력
func dir_read_string(dirpath string) ([]string, []int64) {
	var file_name []string
	var file_info []int64
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// 파일 명(경로포함) 및 파일 크기 출력
			file_name = append(file_name, path)
			file_info = append(file_info, info.Size())
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		//	fmt.Println(list)
		//return files, err
		//	return
	}
	return file_name, file_info
}

func export_csv() {}

func remove_extension(f []string) {
	for _, i := range f {
		d, err := os.Open(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer d.Close()
		fmt.Println(d)
		/*
			files, err := d.Readdir(-1)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(files)

				for _, file := range files {
					if file.Mode().IsRegular() {
						if filepath.Ext(file.Name()) == ".png" {
							os.Remove(file.Name())
							fmt.Println("Deleted", file.Name())
						}
					}
				}*/
	}
}

//func remove_duplicated() {}

func main() {
	var input string
	dirPath := input_path(input)
	file_path, file_info := dir_read_string(dirPath)
	remove_extension(file_path)
	fmt.Println(file_info[0])
	//	fmt.Print(file_path[:], file_info[:])
	//	fileinfo, _ := dir_read_string(dirPath)
	//	fmt.Println(fileinfo)
}
