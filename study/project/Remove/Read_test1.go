//파일 삭제 테스트
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

/*
func DirRead(dir string){
  err :=filepath.Walk(root, walkFn)
}


func DirectoryRead(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}


	for _, f := range files {
		if f.IsDir() {
			fmt.Println(f.Name() + "\t" + "디렉토리")
		} else {
			fmt.Println(f.Name() + "파일" + "\t" + strconv.FormatInt(f.Size(), 10) + "Byte")
		}
	}
}
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	/*	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(runtime.NumCPU()))
		go DirectoryRead("D:/외장하드")
		time.Sleep(5 * time.Second)
	*/
	// 파일 리스트 읽기(디렉토리 포함)
	//	go func() {
	var files []string
	err := filepath.Walk("D:/삭제테스트/외장하드", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		fmt.Println(reflect.TypeOf(info))
		/*	for _, file := range info {
			if file.Mode().IsRegular() {
				if filepath.Ext(file.Name()) == ".docx" {
					fmt.Println(reflect.TypeOf(file))
				}
			}
		}*/
		return nil
	})

	//fmt.Println(path, int(info.Size()))
	if err != nil {
		log.Println(err)
	}
}
