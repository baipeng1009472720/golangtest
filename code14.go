package main

import (
	"fmt"
	"time"
)

var c chan string

func Ping() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprint("ping %d", i)
		i++
	}
}

func test() {
	fmt.Println(<-c)

}
func main() {

	c = make(chan string)

	go Ping()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprint("nihao %d", i)
		fmt.Println(<-c)
	}

	go test()
	c <- "s"
	time.Sleep(2 * time.Second)
}
