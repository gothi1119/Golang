package main

import (
	"flag"
	//"encoding/hex"
	"fmt"
	"godirwalk"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

type fileHash struct {
	path string
	hash uint64
	size int64
	err  error
}

const scanAll = 0
const potentialScanLength = 4096

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
		hashToFiles := make(map[uint64][]fileHash)
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
		fileCount     = 0
		dupCount      = 0
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

	fmt.Printf("%d files, %d duplicates\n", fileCount, dupCount)
}

func getFileChecksum(file *fileHash, scanSize int64) {
	f, err := os.Open(file.path)
	if err != nil {
		file.err = err
		return
	}
	defer f.Close()
	hasher := fnv.New64a()
	if scanSize != scanAll {
		buf := make([]byte, scanSize)
		fmt.Printf("scan first %d bytes of %s\n", scanSize, file.path)
		n, err := f.Read(buf)
		if err == nil {
			hasher.Write(buf[:n])
			file.hash = hasher.Sum64()
			fmt.Printf("%v\n", file.hash)
		} else {
			file.err = err
		}
	} else {
		fmt.Printf("Scanning file %s ...\n", file.path)
		_, file.err = io.Copy(hasher, f)
		if file.err == nil {
			file.hash = hasher.Sum64()
			fmt.Printf("nill %v\n", file.hash)
		}
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage of rmdupes:")
		fmt.Println("rmdupes path : delete duplicate files from [path]")
		fmt.Println("rmdupes path moveto: move duplicate files from [path] to [moveto]")
		return
	}
	var dir = args[0]
	st, err := os.Stat(dir)
	if err != nil {
		panic(err)
	}
	if !st.IsDir() {
		fmt.Println("Invalidate path", dir)
	}
	var moveTo string
	if len(args) > 1 {
		moveTo = args[1]
		if _, err := os.Stat(moveTo); err != nil {
			_ = os.MkdirAll(moveTo, os.ModePerm)
		}
	}

	scanAndRemoveDuplicates(dir, moveTo)
}
