package main

import "fmt"

func main() {
	//
	//a := 1
	//var p *int = &a
	//fmt.Println(*p)
	//
	//if B := 1; B >= 1 {
	//	fmt.Println("ok")
	//}

	//for i := 0; i <= 100; i++ {
	//	fmt.Println(i)
	//}
	//a := 1
	//for {
	//	a++
	//	if a > 3 {
	//		break
	//	}
	//	fmt.Println(a)
	//}
	//
	//b := 1
	//for b <= 3 {
	//	b++
	//	if b > 3 {
	//		break
	//	}
	//	fmt.Println(b)
	//}
	//a := 1
	//switch a {
	//case 0:
	//	fmt.Println("0")
	//case 1:
	//	fmt.Println("1")
	//default:
	//	fmt.Println(a)
	//}

	//switch a := 1; {
	//case a > 0:
	//	fmt.Println("a>0")
	//	fallthrough
	//case a >= 1:
	//	fmt.Println("a>=1")
	//default:
	//	fmt.Println("None")
	//
	//}

LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL
		}
	}

}
