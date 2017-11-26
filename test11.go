package main

import "fmt"

type A struct {
	name string
}

type B struct {
	name string
}

type TZ int

func main() {

	a := A{}
	a.Print()
	fmt.Println(a.name)

	b := B{}
	b.Print()
	fmt.Println(b.name)

	var c TZ
	c.Print()
	fmt.Println(c)
	(*TZ).Print(&c)
}

func (a *A) Print() {
	a.name = "aaa"
	fmt.Println(a.name)
}

func (b B) Print() {
	b.name = "bbb"
	fmt.Println("B")
}

func (a *TZ) Print() {
	*a = 1
	fmt.Println("tz")
}
