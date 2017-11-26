package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name   string
	age    int
	city   string
	id     int
	guitar Guitar
}

type Guitar struct {
	name  string
	brand string
}

func (u User) Hello(name string) {
	fmt.Println("你好，", name, "，我是", u.name)
}
func (g Guitar) get(name string) {
	fmt.Println(g.name, g.brand, name)
}

func main() {

	g := Guitar{"电吉他", "fender"}
	u := User{"baipeng", 21, "北京", 1, g}
	fmt.Println(u)

	//获取对象的类型
	v := reflect.TypeOf(u)
	//获取对象的值
	o := reflect.ValueOf(u)
	fmt.Println(o)
	//检查对象的类型

	if f := v.Kind(); f != reflect.Struct {
		fmt.Println("。。。。", v.Kind())
		return
	}

	//获取对象中属性名称，类型，值
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		val := o.FieldByName(f.Name)
		if f.Type.Kind() == reflect.Struct {
			for j := 0; j <f.Type.NumField(); j++ {
				f2 := f.Type.Field(j)
				val2 := val.FieldByName(f2.Name)
				fmt.Println("??:", f2.Name, f2.Type, val2)
			}

		} else {
			fmt.Println(f.Name, f.Type, val)
		}
	}
	//获取对象包含的方法

	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)
		fmt.Println(m.Name, m.Type, m.Func)

	}
	m := o.MethodByName("Hello")
	args := reflect.Value(reflect.ValueOf("yangqian"))

	fmt.Println(args)
	fmt.Println(m.CanAddr())
	fmt.Println(m.Call([]reflect.Value{reflect.ValueOf("yangqian")}))
}
