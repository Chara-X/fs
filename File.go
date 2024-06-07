package fs

import (
	"io"
	"os"

	"github.com/Chara-X/free"
	"github.com/Chara-X/util/encoding/binary"
)

type File struct {
	file      *os.File
	freePages *free.Heap
	offset    int64
	cur       *Cursor
}

func (f *File) Name() string {
	var buf = make([]byte, binary.ReadInt64At(f.file, f.offset))
	f.file.ReadAt(buf, f.offset+8)
	return string(buf)
}
func (f *File) setName(name string) {
	var buf = []byte(name)
	binary.WriteInt64At(f.file, int64(len(buf)), f.offset)
	f.file.WriteAt(buf, f.offset+8)
}
func (f *File) Size() int64        { return binary.ReadInt64At(f.file, f.offset+24) }
func (f *File) setSize(size int64) { binary.WriteInt64At(f.file, size, f.offset+24) }
func (f *File) page() *Page {
	return &Page{f.file, f.freePages, binary.ReadInt64At(f.file, f.offset+32)}
}
func (f *File) setPage(offset int64) { binary.WriteInt64At(f.file, offset, f.offset+32) }
func (f *File) Read(buf []byte) (n int, err error) {
	if int64(len(buf)) > f.Size()-f.cur.off {
		buf, err = buf[:f.Size()-f.cur.off], io.EOF
	}
	n, _ = f.cur.Read(buf)
	f.Seek(int64(n), 1)
	return
}
func (f *File) Write(buf []byte) (n int, err error) {
	f.setSize(max(f.Size(), f.cur.off+int64(len(buf))))
	n, _ = f.cur.Write(buf)
	f.Seek(int64(n), 1)
	return
}
func (f *File) Seek(off int64, org int) (int64, error) {
	switch org {
	case 0:
		f.cur.off = off
		f.cur.page = f.page()
		f.cur.pageOff = off
	case 1:
		f.cur.off += off
		f.cur.pageOff += off
	case 2:
		return f.Seek(f.Size()-off, 0)
	}
	for f.cur.pageOff >= SizeOfPage-8 {
		f.cur.page = f.cur.page.After()
		f.cur.pageOff -= SizeOfPage - 8
	}
	return f.cur.off, nil
}
