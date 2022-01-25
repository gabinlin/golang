package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	get, err := http.Get("http://localhost:8080")
	if err != nil {
		return
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(all))
}
