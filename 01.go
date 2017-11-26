package main

import (
	"fmt"
	"runtime"
	"os"
)

//我是注释
func add(x int, y int) int {
	return x + y
}

//func person(a int, B int) (int, int) {
//	return a, B
//}

const name = "123"

var v int = 1

type T struct {
	name string
	age  int
}

func init() {

}

const (
	ONE   = 1
	TOW   = 2
	THREE = 3
)

const (
	a = 1
	b = 2
	c = 3
)

type d struct {
	qq  string
	qqw string
}

type (
	str1 string
	age int
)

var (
	nihao=1
	uan = 2
)

func main() {
	println("hello world")

	fmt.Printf("%s\n", runtime.Version())
	println(add(1, 2))
	//println(person(1, 2))
	println(os.Getenv("HOME"))
	println(os.Getenv("USER"))
	println(os.Getenv("GOROOT"))
	a := 1
	println(a)
	str := "123"
	println(str)
	d, f := 1, 2
	println(d, f)

	var aVar = 10
	println(aVar != 5)
	println(aVar != 10)
	aa := uint64(1)

	println(aa)

	var a1 int
	var b1 int32
	a1 = 15
	b1 = b1 + 5
	a1 = a1 + a1

	println(a1, b1)

	println(a, b, c)
	println(nihao)
	println()
}
