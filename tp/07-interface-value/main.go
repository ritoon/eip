package main

import "fmt"

func main() {
	var i interface{}
	i = 7
	fmt.Println("interface value:", i)
	i = "hello"
	fmt.Println("interface value:", i)
	i = true
	fmt.Println("interface value:", i)
}
