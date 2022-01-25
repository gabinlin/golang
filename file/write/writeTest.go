package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("/Users/linjiabin/Downloads/a.txt")
	write, err := file.Write([]byte{88})
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
	fmt.Println(write)

}
