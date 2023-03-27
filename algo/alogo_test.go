package algo

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {

	println("xx")
}

type Stu struct {
}

func TestCache1(t *testing.T) {
	q := make(chan int, 2)
	q <- 1
	q <- 2
	select {
	case q <- 3:
		fmt.Println("ok")
	default:
		for i := 0; i < 3; i++ {
			c := <-q
			fmt.Println(c)
		}

		fmt.Println("wrong")
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	ins := []int{1, 2, 3, 3, 24, 2, 7, 8, 4, 3, 2, 5455, 343, 65464, 54}
	mate := [][]int{ins, ins, ins}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "xxx", args: args{matrix: mate}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			t.Log(tt.args.matrix)
		})
	}
}

func Test_quickMul(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{args: args{x: 2, n: 3}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickMul(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("quickMul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func cew(x float64, n int) float64 {
	ans := 1.0
	arr_ans := x
	for n > 0 {
		if n%2 == 1 {
			ans *= x * arr_ans
		}
		arr_ans *= arr_ans
		n /= 2
	}
	return ans
}

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	nums := []int{-2, 1, -3, 4}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{want: 6, args: args{nums: nums}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{[]int{1, 2, 2, 0, 0, 0, 1, 2, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			t.Log("nums:", tt.args.nums)
		})
	}
}
