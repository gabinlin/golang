package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
	}
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	_, _ = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n"+readString)
	status, err := bufio.NewReader(conn).ReadString('\n')
	err = conn.Close()
	if err != nil {
		return
	}
	fmt.Println(status)
}
