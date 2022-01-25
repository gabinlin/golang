package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copyFile("/Users/linjiabin/Downloads/trainee_second_or_third_stage.bpmn20.xml",
		"/Users/linjiabin/Downloads/trainee_second_or_third_stage.bpmn20.xml_bak")
}

func copyFile(sourcePath string, dstPath string) {
	// 1、read
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(sourceFile)

	// 2、create tmp file
	dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 3、Copy
	writeLen, err := io.Copy(dstFile, sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(writeLen)
}
