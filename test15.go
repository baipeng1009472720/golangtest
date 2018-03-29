package main

import (
	"./test"
	"fmt"
	"time"
)

//idea 会强制提示有返回值
func test2(s []int) []int {
	return append(s, 3)
}
func main() {
	test.Hello()
	test.Init()

	s := make([]int, 0)
	fmt.Println(s)
	s = test2(s)
	fmt.Println(s)

	t := time.Now()
	//时间格式化，不能改常量中的格式
	fmt.Println(t.Format(time.ANSIC))
	//闭包中的坑，变量需要传递
	//s1 := []string{"a", "b", "c"}
	//for _, v := range s1 {
	//	go func() {
	//		fmt.Println(v)
	//	}()
	//}
	//s1 := []string{"a", "b", "c"}
	//for _, v := range s1 {
	//	go func(v string) {
	//		fmt.Println(v)
	//	}(v)
	//}
	//select {}
	var i = 1
	for {
		println(i)
		i++
	}

}
