package algo

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_Unint(t *testing.T) {

	chs := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			chs <- i
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			chs <- i
		}
	}()

	go loop(chs)
	time.Sleep(time.Second * 4)
	close(chs)

}

func loop(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Println("this  value of unbuffer channel", i)
		}
	}
}

func Test_add(t *testing.T) {

	i := new(int)
	*i = 10
	fmt.Println(*i)

	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{want: []int{0, 2}, args: args{k: 4, arr: []int{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := towSum(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
