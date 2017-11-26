package main

import "fmt"

type int2 int

//根据为结构增加方法的知识，尝试声明一个底层类型为int的类型，并实现调用某个方法就递增100
func main() {

	var a int2
	for i := 0; i < 3; i++ {
		a.Increase(100)
		fmt.Println(a)
	}
}

func (a *int2) Increase(num int) {
	*a += int2(num)

}
