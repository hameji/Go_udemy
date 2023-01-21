package main

import "fmt"

func foo(param1, param2 int) {

}

func foo2(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

//
//func main() {
//	foo(10, 20)
//	//	foo(10, 20, 30) // これエラー
//
//	foo2(10, 20)
//	foo2(10, 20, 30)
//	foo2()
//
//	s := []int{1, 2, 3}
//	fmt.Println(s)
//
//	foo2(s...) // 展開していれる、可変長引数
//
//}
