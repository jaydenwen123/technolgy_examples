package fib

import (
	"time"
	"unsafe"
)

//字符串长度
//cpu内存分析
//go test -v -bench . -cpuprofile cpu.out
//go tool pprof cpu.out
//代码覆盖率
//go test -v -bench . -coverprofile cover.out
//go tool cover -html cover.out -o cover.html
func strCount(str string) int {
	var l int
	//第一处：
	//l = len([]rune(str))
	//第一次改进
	//l = len(str)
	//第二次改进
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	runes := *(*[]rune)(unsafe.Pointer(&h))
	l = len(runes)

	return l
}

//计算函数
func algorithm(a int, b int) int {
	time.Sleep(time.Millisecond * 100)
	return a + b
}

type Demo struct {
}

// 注释
//	e.g. t.Test1(123)
func (d *Demo) Test1(v int) int {
	return v + v
}
