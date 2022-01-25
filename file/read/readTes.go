package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1、工具包
	file, _ := os.Open("/Users/linjiabin/Downloads/trainee_second_or_third_stage.bpmn20.xml")
	//all, _ := ioutil.ReadAll(file)
	//fmt.Println(string(all))
	// 2、手写
	var b = make([]byte, 1024, 1024)
	for i := 0; ; i += 1024 {
		at, err := file.ReadAt(b, int64(i))
		if err == io.EOF {
			b = b[0:at]
			fmt.Println(string(b))
			break
		} else {
			fmt.Println(string(b))
		}
	}
	err1 := file.Close()
	if err1 != nil {
		return
	}
}
