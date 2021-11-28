package simple

import "fmt"

type File struct {
	Name string
}

func (f *File) Close() {
	fmt.Println("Closing file", f.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{name}
	return file, func() {
		file.Close()
	}
}
