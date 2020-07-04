package main

import "fmt"

func main() {
	//按位取反的公式：~a=-(a+1)
	var a = 5
	fmt.Println(a)
	var b = ^a
	//按位取反
	fmt.Println(b)
	//按位取反
	fmt.Println(^b)
}
