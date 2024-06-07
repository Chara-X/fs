package fs

import (
	"os"

	"github.com/Chara-X/free"
)

type FS struct {
	file      *os.File
	files     map[string]int64
	pages     map[int64]struct{}
	freeFiles *free.Heap
	freePages *free.Heap
}

func New(name string) *FS {
	var file, _ = os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if stat, _ := file.Stat(); stat.Size() == 0 {
		file.Write(make([]byte, SizeOfFS))
	}
	var fs = &FS{file, map[string]int64{}, map[int64]struct{}{}, new(free.Heap), new(free.Heap)}
	var offset int64
	for ; offset < SizeOfFiles; offset += SizeOfFile {
		var file = &File{file, fs.freePages, offset, new(Cursor)}
		if len(file.Name()) > 0 {
			fs.files[file.Name()] = file.offset
			for page := file.page(); page.offset != 0; page = page.After() {
				fs.pages[page.offset] = struct{}{}
			}
		} else {
			fs.freeFiles.Push(offset)
		}
	}
	for offset = SizeOfFiles; offset < SizeOfFS; offset += SizeOfPage {
		if _, ok := fs.pages[offset]; !ok {
			fs.freePages.Push(offset)
		}
	}
	return fs
}
func (f *FS) Open(name string) *File {
	var file = &File{f.file, f.freePages, f.files[name], new(Cursor)}
	if _, ok := f.files[name]; !ok {
		file.offset = f.freeFiles.Pop()
		file.setName(name)
		file.setPage(f.freePages.Pop())
		f.files[name] = file.offset
	}
	file.cur.page = file.page()
	return file
}
func (f *FS) Remove(name string) {
	var file = f.Open(name)
	file.setName("")
	delete(f.files, name)
	f.freeFiles.Push(file.offset)
	for page := file.page(); page.offset != 0; page = page.After() {
		f.freePages.Push(page.offset)
	}
}
func (f *FS) Close() { f.file.Close() }
