package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			// handle err2
			continue
		}
		var b = make([]byte, 1024, 1024)
		_, err2 := conn.Read(b)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println(string(b))
		_, err = conn.Write(b)
		if err != nil {
			return
		}
		err = conn.Close()
		if err != nil {
			return
		}
	}

}
