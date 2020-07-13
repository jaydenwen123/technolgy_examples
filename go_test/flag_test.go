package fib

import (
	"flag"
	"testing"
)

var flag1 string

var flag2 int

func init() {
	flag.StringVar(&flag1,"var1","hello","enter var1")
	flag.IntVar(&flag2,"var2",3,"enter var2")
}

func TestFlag(t *testing.T){
	t.Logf("flag1:%v",flag1)
	t.Logf("flag2:%v",flag2)
}



func TestArgs(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}

	argList := flag.Args() // flag.Args() 返回 -args 后面的所有参数，以切片表示，每个元素代表一个参数
	t.Log("argList:",argList)
	for _, arg := range argList {
		if arg == "cloud" {
			t.Log("Running in cloud.")
		}else {
			t.Log("Running in other mode.")
		}
	}
}
