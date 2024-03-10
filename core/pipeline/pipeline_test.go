package pipeline_test

import (
	"context"
	"fmt"
	"github.com/SeeJson/node/core/pipeline"
	"sort"
	"testing"
	"time"
)

func ExampleMergeChannels() {
	inChan1 := make(chan interface{}, 10)
	inChan2 := make(chan interface{}, 10)

	inChan1 <- 1
	inChan1 <- 3
	inChan1 <- 5
	inChan1 <- 7
	inChan1 <- 9
	close(inChan1)

	inChan2 <- 2
	inChan2 <- 4
	inChan2 <- 6
	inChan2 <- 8
	inChan2 <- 10
	close(inChan2)

	outChan := pipeline.MergeChannels([]chan interface{}{inChan1, inChan2})

	var ints []int
	for e := range outChan {
		ints = append(ints, e.(int))
	}
	sort.Ints(ints)
	fmt.Println(ints)

	// Output: [1 2 3 4 5 6 7 8 9 10]
}

type Author struct {
	Name         int      `json:Name,tag:"22"`
	Publications []string `json:Publication,omitempty`
}

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("Eat")
}

var word = make(chan struct{}, 1)
var num = make(chan struct{}, 1)

func printNums() {
	for i := 0; i < 10; i++ {
		<-word
		fmt.Println(1)
		num <- struct{}{}
	}
}

func printWord() {
	for i := 0; i < 10; i++ {
		<-num
		fmt.Println("a")
		word <- struct{}{}
	}
}

/*
两个协程交替打印10个字母和数字
*/
func TestMergeChannels(t *testing.T) {
	num <- struct{}{}
	go printNums()
	go printWord()

	time.Sleep(time.Second * 1)
	fmt.Println("ssds")
}

/*
启动 2个groutine 2秒后取消， 第一个协程1秒执行完，第二个协程3秒执行完。
*/
func f1(in chan struct{}) {
	time.Sleep(time.Second * 1)
	in <- struct{}{}
}

func f2(in chan struct{}) {
	time.Sleep(time.Second * 3)
	in <- struct{}{}
}

func Test2s(t *testing.T) {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	go func() {
		go f1(ch1)
		select {
		case <-ch1:
			fmt.Println("f1 done")
		case <-ctx.Done():
			fmt.Println("f1 timeout")
		}
	}()

	go func() {
		go f2(ch2)
		select {
		case <-ch2:
			fmt.Println("f2 done")
		case <-ctx.Done():
			fmt.Println("f2 timeout")
		}
	}()

	time.Sleep(5 * time.Second)
}

/*
当select监控多个chan同时到达就绪态时，如何先执行某个任务？
*/
func priority_select(ch1, ch2 <-chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		//default:
		//	time.Sleep(1 * time.Second)
		//	fmt.Println("end")
		case val2 := <-ch2:
		priority:
			for {
				select {
				case val1 := <-ch1:
					fmt.Println(val1)

				default:
					break priority
				}
			}
			fmt.Println(val2)
		}
	}

}
func Test3(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go priority_select(ch1, ch2)
	ch2 <- "2"
	ch1 <- "1"
	time.Sleep(1 * time.Second)
	ch2 <- "2"
	ch1 <- "1"

	time.Sleep(10 * time.Second)
}

func TestCheck1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
