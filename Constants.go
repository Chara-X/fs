package fs

const (
	NumberOfFiles = 1 << 3
	NumberOfPages = 1 << 6
	SizeOfFile    = 8 + 16 + 8 + 8
	SizeOfPage    = 8 + 8
	SizeOfFiles   = NumberOfFiles * SizeOfFile
	SizeOfPages   = NumberOfPages * SizeOfPage
	SizeOfFS      = SizeOfFiles + SizeOfPages
)
