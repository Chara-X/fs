package fs

import (
	"fmt"
	"io"
)

func Dump(fs *FS) {
	fmt.Println("Number of files:", len(fs.files))
	fmt.Println("Number of free files:", fs.freeFiles.Len())
	fmt.Println("Number of free pages:", fs.freePages.Len())
	for name := range fs.files {
		var data, _ = io.ReadAll(fs.Open(name))
		fmt.Println("-", len(name), ":", name)
		fmt.Println(" ", len(data), ":", string(data))
	}
}
