package main

import "fmt"

func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("def i = ", i)
		defer func() { fmt.Println("def  clo i = ", i) }()
		fs[i] = func() {
			fmt.Println("clo i = ", i)
		}
	}
	for _,f := range fs{
		  f()
	}
}
