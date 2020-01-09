package main

import (

	//	"encoding/hex"
	"crypto/sha1"
	"flag"
	"fmt"
	"godirwalk"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

type fileHash struct {
	path string
	hash string
	size int64
	err  error
}

const scanAll = 0
const potentialScanLength = 4096

//대상 디렉터리 경로를 입력받음
func input_path(dir string) string {
	fmt.Println("경로 입력:")
	fmt.Scan(&dir)
	return dir
}

/*
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
	}
	return file_name, file_info
}

//func export_csv() {}

/* func remove_extension(f []string) { //Plz, Check this
	//파일 목록이 저장된 슬라이스 입력 및 i에 리스트 저장
	for _, i := range f {
		file, err := os.Stat(i)
		if err != nil {
			panic(err)
		}
		if filepath.Ext(file.Name()) == ".png" {
			os.Remove(file.Name())
			fmt.Println("Deleted", file.Name())
		}
	}
}
*/
/*
func remove_duplicated(f []string) {
	fileMap := make(map[string][]string)
	//	var ret1 string
	h := sha1.New()
	for _, i := range f {
		t, err := os.Open(i)
		if err != nil {
			panic(err)
		}
		defer t.Close()
		if _, err := io.Copy(h, t); err != nil {
			panic(err)
		}

		hashBytes := h.Sum(nil)
		ret := hex.EncodeToString(hashBytes)
		fileMap[ret] = append(fileMap[ret], i)
	}
	//	fmt.Println(fileMap)
	for _, item := range fileMap {
		fmt.Println(item)
		//firstFile := item[0]
		for i := 1; i < len(item); i++ {
			file := item[i]
			fmt.Println(file)
		}

	}
}*/
func getDuplicates(potentialDups [][]fileHash, scanLength int64) [][]fileHash {
	maxFds := runtime.NumCPU()
	throttle := make(chan bool, maxFds)
	for _, files := range potentialDups {
		for idx := 0; idx < len(files); idx++ {
			if scanLength != scanAll || files[idx].size > potentialScanLength {
				throttle <- true
				go func(p *fileHash) {
					getFileChecksum(p, scanLength)
					<-throttle
				}(&files[idx])
			}
		}
	}
	for i := 0; i < maxFds; i++ {
		// Will block until goroutines have removed all entries they have put in.
		throttle <- true
	}
	duplicates := make([][]fileHash, 0, len(potentialDups))
	for _, files := range potentialDups {
		hashToFiles := make(map[string][]fileHash)
		for _, file := range files {
			if file.err == nil {
				files, ok := hashToFiles[file.hash]
				if !ok {
					files = make([]fileHash, 0, 2)
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

func removeDuplicates(duplicates [][]fileHash, moveDuplicateTo string) (dupCount int) {
	for _, files := range duplicates {
		fmt.Println("Original is", files[0].path)
		for _, path := range files[1:] {
			dupCount += 1
			fmt.Println("deleting", path.path)
			if moveDuplicateTo != "" {
				fileIndex := 0
				for {
					var filename string
					if fileIndex == 0 {
						filename = filepath.Join(moveDuplicateTo, filepath.Base(path.path))
					} else {
						filename = filepath.Base(path.path)
						ext := filepath.Ext(path.path)
						name := filename[0 : len(filename)-len(ext)]
						filename = filepath.Join(moveDuplicateTo, fmt.Sprintf("%s(%d)%s", name, fileIndex, ext))
					}

					if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
						err := os.Rename(path.path, filename)
						if err != nil {
							fmt.Println(err)
						}
						break
					}
					fileIndex += 1
				}

			} else {
				_ = os.Remove(path.path)
			}
		}

	}
	return dupCount
}

func scanAndRemoveDuplicates(root string, moveDuplicateTo string) {
	var (
		fileCount = 0
		//	dupCount      = 0
		sameSizeCount = 0
		fileByeSize   = make(map[int64][]string)
	)
	err := godirwalk.Walk(root, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {
			if de.IsDir() {
				return nil
			}
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

	sameSizeFiles := make([][]fileHash, 0, sameSizeCount)
	for size, files := range fileByeSize {
		if len(files) > 1 {
			fh := make([]fileHash, len(files))
			for i := 0; i < len(files); i++ {
				fh[i] = fileHash{path: files[i], size: size}
			}
			sameSizeFiles = append(sameSizeFiles, fh)
		}
	}

	if len(sameSizeFiles) > 0 {
		potentialDups := getDuplicates(sameSizeFiles, potentialScanLength)
		if len(potentialDups) > 0 {
			duplicates := getDuplicates(potentialDups, scanAll)
			if len(duplicates) > 0 {
				dupCount = removeDuplicates(duplicates, moveDuplicateTo)
			}
		}
	}
	fmt.Printf("%v files, %d duplicates\n", sameSizeFiles, dupCount)
}

func getFileChecksum(file *fileHash, scanSize int64) {
	f, err := os.Open(file.path)
	if err != nil {
		file.err = err
		return
	}
	defer f.Close()
	hasher := sha1.New()
	if scanSize != scanAll {
		buf := make([]byte, scanSize)
		fmt.Printf("scan first %d bytes of %s\n", scanSize, file.path)
		n, err := f.Read(buf)
		if err == nil {
			hasher.Write(buf[:n])
			file.hash = string(hasher.Sum(nil))

		} else {
			file.err = err
		}
	} else {
		fmt.Printf("Scanning file %s ...\n", file.path)
		_, file.err = io.Copy(hasher, f)
		if file.err == nil {
			file.hash = string(hasher.Sum(nil))
		}
	}
}

func main() {
	var input string
	args := flag.Args()
	dirPath := input_path(input)
	//	file_path, _ /*file_info*/ := dir_read_string(dirPath)
	//	remove_extension(file_path)
	//fmt.Println(file_info)
	st, err := os.Stat(dirPath)
	if err != nil {
		panic(err)
	}
	if !st.IsDir() {
		fmt.Println("Invaild Path", dirPath)
	}
	var moveTo string
	if len(args) > 1 {
		moveTo = args[1]
		if _, err := os.Stat(moveTo); err != nil {
			_ = os.MkdirAll(moveTo, os.ModePerm)
		}
	}
	//remove_duplicated(dirPath)
	scanAndRemoveDuplicates(dirPath, moveTo)
}
