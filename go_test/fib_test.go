package fib

import (
	"flag"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "n=5",
			args: args{
				n: 5,
				//1,1,2,3,5
			},
			want: 8,
		},
		{
			name: "n=8",
			args: args{
				n: 8,
			},
			want: 34,
		}, {
			name: "n=12",
			args: args{
				n: 12,
			},
			want: 233,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFib2(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	t.Logf("the flag args:%v", flag.Args())
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "n=5",
			args: args{
				n: 5,
				//1,1,2,3,5
			},
			want: 8,
		},
		{
			name: "n=8",
			args: args{
				n: 8,
			},
			want: 34,
		}, {
			name: "n=12",
			args: args{
				n: 12,
			},
			want: 233,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib2(tt.args.n); got != tt.want {
				t.Errorf("Fib2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	b.ResetTimer()
	n := 20
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib2(b *testing.B) {
	//b.ResetTimer()
	n := 20
	for i := 0; i < b.N; i++ {
		Fib2(n)
	}
}

