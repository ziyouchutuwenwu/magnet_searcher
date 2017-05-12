package file

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	jsonFile := Create()
	jsonFile.Init()
	jsonFile.SetFileName("./file.go")

	if jsonFile.IsFileExist() {
		jsonFile.Open()
		fmt.Println("文件存在")
		content := jsonFile.ReadFileString()
		t.Log(content)
	} else {
		fmt.Println("文件不存在")
	}
}
