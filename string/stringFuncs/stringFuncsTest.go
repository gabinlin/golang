package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 方法1
	var str = "你好吗aa"
	for _, i := range str {
		fmt.Printf("%c", i)
	}
	fmt.Println("\n**************************")
	// 方法2
	var arr = []rune(str)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%c", arr[i])
	}
	fmt.Println("\n**************************")

	var a = strconv.Itoa(98)
	fmt.Println(a)

	fmt.Printf("%c\n", 98)

	fmt.Println(strings.Contains(" Hello", "Hello"))

	fmt.Println(strings.EqualFold("hello", "HELLo"))

	var b = "hello"
	var c = "hello"
	fmt.Println(b == c)

	fmt.Println(strings.Index("Hello", "o"))

	fmt.Println(strings.Replace("haha", "ha", "golang", -1))

	upper := strings.ToUpper("golang")
	fmt.Printf("%v\n", upper)
	fmt.Printf("%v\n", strings.ToLower(upper))

	fmt.Println(strings.Split("he,he", ","))

	fmt.Println(strings.TrimSpace(" hahahha  a  "))

	fmt.Println(strings.Trim("aaaa哈哈aaaa", "a"))

	fmt.Println(strings.HasSuffix("DC_201", "201"))
	fmt.Println(strings.HasPrefix("DC_201", "DC"))
}
