package main

import (
	"fmt"
	"godirwalk"
	"hash/fnv"
	"io"
	"path/filepath"
	"runtime"
	//"io"

	"os"
)

type filehash struct {
	path string
	hash uint64
	size int64
	err  error
}

const scanAll = 0
const ScanLength = 8192

//대상 디렉터리 경로를 입력받음
func input_path(dir string) string {
	fmt.Println("경로 입력:")
	fmt.Scan(&dir)
	return dir
}

func scanDir(root string) {
	var (
		fileCount     = 0
		dupCount      = 0
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

	samesizeFiles := make([][]filehash, 0, sameSizeCount)
	for size, files := range fileByeSize {
		if len(files) > 1 {
			fh := make([]filehash, len(files))
			for i := 0; i < len(files); i++ {
				fh[i] = filehash{path: files[i], size: size}

			}
			samesizeFiles = append(samesizeFiles, fh)
		}

	}
	fileRemoveExt(file_list)
	if len(samesizeFiles) > 0 {
		potentialDups := getDuplicates(samesizeFiles, ScanLength)
		if len(potentialDups) > 0 {
			duplicates := getDuplicates(potentialDups, scanAll)
			if len(duplicates) > 0 {
				dupCount = removeDuplicates(duplicates)
			}
		}
	}
	fmt.Printf("%d files, %d duplicates\n", fileCount, dupCount)
}

func getDuplicates(potentialDups [][]filehash, scanLength int64) [][]filehash {
	runtime.GOMAXPROCS(runtime.NumCPU())
	maxFds := runtime.NumCPU()
	throttle := make(chan bool, maxFds)
	for _, files := range potentialDups {
		for idx := 0; idx < len(files); idx++ {
			if scanLength != scanAll || files[idx].size > ScanLength {
				throttle <- true
				go func(p *filehash) {
					getFileChecksum(p, scanLength)
					<-throttle
				}(&files[idx])
			}
		}
	}
	for i := 0; i < maxFds; i++ {
		throttle <- true
	}
	duplicates := make([][]filehash, 0, len(potentialDups))
	for _, files := range potentialDups {
		hashToFiles := make(map[uint64][]filehash)
		for _, file := range files {
			if file.err == nil {
				files, ok := hashToFiles[file.hash]
				if !ok {
					files = make([]filehash, 0, 2)
				}
				hashToFiles[file.hash] = append(files, file)
			}
		}
		for _, files := range hashToFiles {
			if len(files) > 1 {
				duplicates = append(duplicates, files)
			}
		}
	}
	return duplicates
}
func removeDuplicates(duplicates [][]filehash) (dupCount int) {
	for _, files := range duplicates {
		fmt.Println("Original is", files[0].path)
		for _, path := range files[1:] {
			dupCount += 1
			fmt.Println("deleting", path.path)
			_ = os.Remove(path.path)
		}
	}
	return dupCount
}
func getFileChecksum(file *filehash, scanSize int64) {
	list, err := os.Open(file.path)
	if err != nil {
		file.err = err
		return
	}
	defer list.Close()
	file_hash := fnv.New64a()
	if scanSize != scanAll {
		buf := make([]byte, scanSize)
		fmt.Printf("scan first %d bytes of %s\n", scanSize, file.path)
		n, err := list.Read(buf)
		if err == nil {
			file_hash.Write(buf[:n])
			file.hash = file_hash.Sum64()
			//	fmt.Printf("\n", file.hash)
		} else {
			file.err = err
		}
	} else {
		fmt.Printf("Scanning file %s...\n", file.path)
		_, file.err = io.Copy(file_hash, list)
		if file.err == nil {
			file.hash = file_hash.Sum64()
		}
	}
}

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
