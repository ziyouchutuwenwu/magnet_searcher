package file

import (
	"io"
	"os"
	"io/ioutil"
)

type File struct {
	fileName string
	file     *os.File
}

func Create() *File {
	file := new(File)
	return file
}

func (this *File) Init() {
	this.fileName = ""
	this.file = nil
}

func (this *File) SetFileName(fileName string) {
	this.fileName = fileName
}

func (this *File) IsFileExist() bool {
	var _exist = true

	if 0 == len(this.fileName) {
		_exist = false
	}
	if _, err := os.Stat(this.fileName); os.IsNotExist(err) {
		_exist = false
	}

	return _exist
}

func (this *File) Open() {
	this.file, _ = os.OpenFile(this.fileName, os.O_APPEND, 0666)
}

func (this *File) Close() {
	this.file.Close()
}

func (this *File) Create() {
	this.file, _ = os.Create(this.fileName)
}

func (this *File) ReadFileByte() []byte {
	content, _ := ioutil.ReadAll(this.file)
	return content
}

func (this *File) ReadFileString() string {
	content := this.ReadFileByte()
	return string(content)
}

func (this *File) WriteFileString(content string) {
	io.WriteString(this.file, string(content))
}

func (this *File) GetFileSize() int64 {
	f, e := os.Stat(this.fileName)
	if e != nil {
		return 0
	}
	return f.Size()
}
