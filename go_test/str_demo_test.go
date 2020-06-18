package fib

import (
	"fmt"
	"testing"
)

func Test_strCount(t *testing.T) {
	type args struct {
		str string
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{"name1", args{"阿杜zhenxun"}, 9},
		{"name2", args{"ado"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strCount(tt.args.str); got != tt.want {
				t.Errorf("strCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleDemo_Test1() {
	fmt.Println((&Demo{}).Test1(1))

	// Output:2
}

func Benchmark_algorigthm(b *testing.B) {
	args1 := 1
	args2 := 2
	want := 3
	b.Run("name1", func(bb *testing.B) {
		for i := 0; i < 10; i++ {
			if got := algorithm(args1, args2); got != want {
				b.Errorf("algorithm() = %v, want %v", got, want)
			}
		}
	})
}

func Benchmark_strCount(b *testing.B){
	s:="阿杜ado"
	for i:=0;i<25;i++{
		s = s+s
	}
	//b.Logf("len(s) = %d",len(s))
	b.ResetTimer() //这里将生成字符串的时间去除
	for i:=0;i<b.N;i++{
		strCount(s)
	}
}
