package fs

type Cursor struct {
	off     int64
	page    *Page
	pageOff int64
}

func (c *Cursor) Read(buf []byte) (n int, err error)  { return c.page.ReadAt(buf, c.pageOff) }
func (c *Cursor) Write(buf []byte) (n int, err error) { return c.page.WriteAt(buf, c.pageOff) }
