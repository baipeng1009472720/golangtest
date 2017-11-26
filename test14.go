package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

//并发 channel 通过通道的机制通信

func main() {

	t := time.Now()
	fmt.Println(t)
	////有缓存，异步的不阻塞
	//c := make(chan bool, 1)
	//go func() {
	//	fmt.Println("Go Go Go!!")
	//	<-c
	//
	//	//close(c)
	//}()
	//c <- true
	//
	////for v := range c {
	////	fmt.Println(v)
	////}
	//
	////无缓存，同步阻塞
	//c1 := make(chan bool)
	//go func() {
	//	fmt.Println("Go Go Go!!")
	//	<-c1
	//}()
	//c1 <- true

	//
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait()
	fmt.Println(time.Now())

	//time.Sleep(1 * time.Second)

	//一个或者多个通道的处理 select
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "baipeng"
	c1 <- 2
	c2 <- "yangqian"

	close(c1)
	//close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}

	c3 := make(chan int)
	go func() {
		for v := range c3 {
			fmt.Println(v)
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case c3 <- 0:
		case c3 <- 1:
		}
	}

	c4 := make(chan bool)
	select {
	case v := <-c4:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("time")

	}

}

func Go(wg *sync.WaitGroup, i int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(i, a)
	//if i == 9 {
	wg.Done()
	//close(c)
	//}
}
