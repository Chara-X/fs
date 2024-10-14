# File system in one file

```go
func ExampleFS() {
	var f = fs.New("fs.db")
	defer f.Close()
	// Example 1: Create two files.
	var f1 = f.Open("file1.txt")
	var f2 = f.Open("file2.txt")
	f1.Write([]byte("I'm a file! Named file1.txt."))
	f2.Write([]byte("I'm a file! And i named file2.txt."))
	fs.Dump(f)
	// Example 2: Remove a file.
	f.Remove("file1.txt")
	fs.Dump(f)
	// Example 3: Create a new file.
	var f3 = f.Open("file3.txt")
	f3.Write([]byte("I'm a new file! Named file3.txt."))
	fs.Dump(f)
	// Example 4: Append data to a file.
	var fx = f.Open("file2.txt")
	fx.Seek(0, 2)
	fx.Write([]byte(" I'm appended data!"))
	fs.Dump(f)
}
```

## Reference

[Case Study: A Simple Filesystem](https://azrael.digipen.edu/~mmead/www/Courses/CS180/Simple-FS.html)
