package main

import (
	"fmt"
	"godirwalk"
	"path/filepath"
	//"io"
	"os"
	//"runtime"
	//"path/filepath"
	_ "crypto/sha1"
)

type filehash struct {
	path string
	hash byte
	size int64
	err  error
}

const scanAll = 0
const ScanLength = 4096

//대상 디렉터리 경로를 입력받음
func input_path(dir string) string {
	fmt.Println("경로 입력:")
	fmt.Scan(&dir)
	return dir
}

func scanDir(root string) {
	var (
		fileCount = 0
		//dupCount      = 0
		sameSizeCount = 0
		fileByeSize   = make(map[int64][]string)
		file_list     []string
	)

	err := godirwalk.Walk(root, &godirwalk.Options{
		Callback: func(path string, ph *godirwalk.Dirent) error {
			if ph.IsDir() {
				return nil
			}
			file_list = append(file_list, path)
			info, err := os.Stat(path)
			if err == nil {
				fileCount += 1
				size := info.Size()
				if size > 0 {
					files, ok := fileByeSize[size]
					if !ok {
						files = make([]string, 0, 2)
					} else {
						sameSizeCount += 1
					}
					fileByeSize[size] = append(files, path)
				}
			} else {
				fmt.Println(err)
			}
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
		Unsorted: true,
	})
	if err != nil {
		panic(err)
	}

	//	filelist := make([][]filehash, 0, sameSizeCount)
	for size, files := range fileByeSize {
		if len(files) > 1 {
			fh := make([]filehash, len(files))
			for i := 0; i < len(files); i++ {
				fh[i] = filehash{path: files[i], size: size}

			}
		}

	}
	fileRemoveExt(file_list)
}
func getFileChecksum(file *fileHash, scanSize int64)

//fileRemoveExt(files []string)
//fileRemoveDuplicates()
//}
func fileRemoveExt(filelist []string) {
	for _, files := range filelist {
		file, err := os.Stat(files)
		if err != nil {
			panic(err)
		}
		fileName := file.Name()
		if filepath.Ext(fileName) == ".bmp" || filepath.Ext(fileName) == ".BMP" {
			os.Remove(files)

		}
	}
}

func main() {
	var input string
	dirPath := input_path(input)
	st, err := os.Stat(dirPath)
	if err != nil {
		panic(err)
	}
	if !st.IsDir() {
		fmt.Println("Invaild Path", dirPath)
	}
	scanDir(dirPath)
}
