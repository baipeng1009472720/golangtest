package main

import (
	"fmt"
)

//机构struct 10 课时
//注意值传递，不会改变原值

//使用寻址，传递
type person struct {
	name string
	age  int
}

type person2 struct {
	name string
	age  int
	contact struct {
		phone, city string
	}
}

type person3 struct {
	string
	int
}

type test1 struct {
	name string
}
type test2 struct {
	name string
}

type human struct {
	sex int
}
type teacher struct {
	human
	name string
	age  int
}
type student struct {
	name string
	age  int
	human
}



func main() {
	a := &person{"baipeng", 25}
	fmt.Println(a.age)
	fmt.Println(a.name)
	a.name = "yangqian"
	fmt.Println(a.name)
	lists := []person{}
	for i := 0; i < 10; i++ {
		lists = append(lists, person{"baipeng", i})
	}

	fmt.Println(lists)
	fmt.Println(len(lists))
	fmt.Println(cap(lists))
	A(a)
	B(a)
	fmt.Println(a.name)

	//匿名结构
	b := struct {
		name string
		age  int
	}{name: "baipeng", age: 25}

	fmt.Println(b)

	//匿名结构嵌套其他结构

	c := person2{name: "xiaobai", age: 25}
	c.contact.city = "北京"
	c.contact.phone = "18511982106"
	fmt.Println(c)

	//匿名字段
	d := person3{"123", 2}
	fmt.Println(d)

	t1 := test1{"aaa"}
	t2 := test1{"aaa"}

	fmt.Println(t1 == t2)

	//类似于继承的属性复用

	s1 := teacher{name: "baipeng", age: 21, human: human{sex: 1}}
	s2 := teacher{name: "yangqian", age: 21, human: human{sex: 1}}
	fmt.Println(s1)
	fmt.Println(s2)
	s1.human = human{sex: 2}
	fmt.Println(s1)
//	http:gowalker.org
}

func A(per *person) {
	per.name = "xiuwen"
	fmt.Println(per)
}

func B(per *person) {
	per.name = "baipeng123"
	fmt.Println(per)
}
