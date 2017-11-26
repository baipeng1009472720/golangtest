package main

import "fmt"

//12 接口interface

type USB interface {
	Name() string
	Connecter
}

type PhoneConnect struct {
	name string
}

type PhoneConnect2 struct {
	name string
}

func (pc PhoneConnect2) Name() string {
	return pc.name
}

func (pc PhoneConnect2) Connect() {
	fmt.Println("connect:", pc.name)
}

func (pc PhoneConnect) Name() string {
	return pc.name
}

func (pc PhoneConnect) Connect() {
	fmt.Println("connect:", pc.name)
}

//嵌入接口

type Connecter interface {
	Connect()
}

type TVConnecte struct {
	name string
}

func (tv TVConnecte) Name() string {
	return tv.name
}

func (tv TVConnecte) Connect() {
	fmt.Println("TVConnecte", tv.name)
}
func main() {

	//var a USB
	//a = PhoneConnect{name: "baipeng"}
	//a.Connect()
	//
	//Disconneect(a)
	//Disconneect2(PhoneConnect2{name: "123"})
	//
	//tv := TVConnecte{"tv"}
	//a = USB(tv)
	//a.Connect()

	pc := PhoneConnect{"asads"}
	var b Connecter
	b = Connecter(pc)
	b.Connect()
	pc.name = "pc"
	b.Connect()

	var a interface{}
	fmt.Println(a)
	fmt.Println(a == nil)

	var p *int = nil
	a = p
	fmt.Println(a)
	fmt.Println(a == nil)
}

func Disconneect(usb USB) {
	//类型判断
	if pc, ok := usb.(PhoneConnect); ok {
		fmt.Println("Disconneect:", pc.Name())
		return
	}
	fmt.Println("devive")
}

//空接口传值

func Disconneect2(usb interface{}) {
	//类型判断
	switch v := usb.(type) {

	case PhoneConnect:
		fmt.Println("PhoneConnect:", v.Name())
	case PhoneConnect2:
		fmt.Println("PhoneConnect2:", v.Name())

	default:
		fmt.Println("devive2")

	}
}

//接口是一个活多个方法签名的集合
