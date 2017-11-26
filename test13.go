package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {

	fmt.Println("Hello", name, "my mane is ", u.Name)
}

type Manager struct {
	User
	title string
}

//课时13 反射机制reflection
func main() {

	u := User{1, "ok", 12}
	Info(u)

	m := Manager{u, "213"}
	fmt.Println(m)
	t := reflect.TypeOf(m)
	fmt.Println(t.FieldByName("Id"))
	fmt.Println(t.FieldByIndex([]int{0}))

	//反射修改对象的值
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

	//
	u2 := User{2, "baipeng", 21}
	Set(&u2)
	fmt.Println(u2)
	//方法的调用
	u2.Hello("baiee")

	v2 := reflect.ValueOf(u2)

	mv := v2.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("joe")}

	f := mv.Call(args)
	fmt.Println(f)
}

func Set(o interface{}) {

	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name");
	if !f.IsValid() {
		fmt.Println("bad")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("yangqian")
	}
}

func Info(o interface{}) {

	t := reflect.TypeOf(o)

	//对类型的name
	fmt.Println("type :", t.Name())

	v := reflect.ValueOf(o)

	//对象的值
	fmt.Println("fields:,", v)

	//对象中每个属性 类型 和值
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
