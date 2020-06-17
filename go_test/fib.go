package fib

import "fmt"

func Fib(n int) int {
	if n <= 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Fib2(n int) int {
	if n <= 2 {
		return n
	}
	first, second := 1, 2
	var tmp int
	var mem []string
	mem = make([]string, 1000)
	for i := 3; i <= n; i++ {
		tmp = first
		first = second
		//分配内存
		second = tmp + second
		mem = append(mem, fmt.Sprintf("%d", second))
	}
	return second
}
