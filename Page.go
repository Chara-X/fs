package fs

import (
	"os"

	"github.com/Chara-X/free"
	"github.com/Chara-X/util/encoding/binary"
)

type Page struct {
	file      *os.File
	freePages *free.Heap
	offset    int64
}

func (p *Page) After() *Page          { return &Page{p.file, p.freePages, binary.ReadInt64At(p.file, p.offset)} }
func (p *Page) SetAfter(offset int64) { binary.WriteInt64At(p.file, offset, p.offset) }
func (p *Page) ReadAt(buf []byte, off int64) (n int, err error) {
	n = len(buf)
	if n, _ := p.file.ReadAt(buf[:min(int64(n), SizeOfPage-8-off)], p.offset+8+off); len(buf[n:]) > 0 {
		buf = buf[n:]
		p.After().ReadAt(buf, 0)
	}
	return
}
func (p *Page) WriteAt(buf []byte, off int64) (n int, err error) {
	n = len(buf)
	if n, _ := p.file.WriteAt(buf[:min(int64(n), SizeOfPage-8-off)], p.offset+8+off); len(buf[n:]) > 0 {
		buf = buf[n:]
		if p.After().offset == 0 {
			p.SetAfter(p.freePages.Pop())
		}
		p.After().WriteAt(buf, 0)
	}
	return
}
