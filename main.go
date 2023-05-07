package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Wb struct {
	Obj *int
}

var A Wb
var B Wb

func simpleSet(c *int) {
	A.Obj = nil
	B.Obj = c
	A.Obj = c
	B.Obj = nil
	fmt.Println("A:%d,b:%d", A, B)
}

type Value struct {
	Name   string
	GoodAt []string
	map1   map[string]string
}

/*func main() {
	ctx, cancel := context.WithCancel(context.Background())
	num := 0
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine exit")
				return
			case <-time.After(time.Second):
				num++
				fmt.Printf("goroutine wait times: %d\n", num)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 4)
	cancel()
	fmt.Println("main exit")
	fmt.Printf("goroutine wait times: %d\n", num)
}
*/

var (
	counter int32          //计数器
	wg      sync.WaitGroup //信号量
)

var ch chan struct{} = make(chan struct{}, 2)

func foo() {
	ch <- struct{}{}
	log.Println("foo() 000")
	ch <- struct{}{}
	log.Println("foo() 111")
	time.Sleep(5 * time.Second)
	log.Println("foo() 222")
	close(ch)
	log.Println("foo() 333")
}

func main() {
	var b struct{}

	log.Println("main() 111")
	go foo()
	log.Println("main() 222")
	a := <-ch
	log.Println("main() 333", a)
	b = <-ch
	log.Println("main() 444", b)
	c := <-ch
	log.Println("main() 555", c)
}

func incCounter(index int) {
	defer wg.Done()
	spinNum := 0
	for {
		//2.1原子操作
		old := counter
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	fmt.Printf("thread,%d,spinnum,%d\n", index, spinNum)
}
