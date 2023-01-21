package main

import (
	"fmt"
)

func boo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

// func main() {
// 	//boo()
// 	//defer fmt.Println("world") // 関数が終わった後に実行される
// 	//
// 	//fmt.Println("Hello")
// 	//
// 	//fmt.Println("run")
// 	//defer fmt.Println(1)
// 	//defer fmt.Println(2)
// 	//defer fmt.Println(3)
// 	//// スタッキングデファー
// 	//// はじめに入れたものが最後に実行される
// 	//
// 	fmt.Println("success")

// 	// file, _ := os.Open("./lesson.go")
// 	// defer file.Close()
// 	// data := make([]byte, 100)
// 	// file.Read(data)
// 	// fmt.Println(string(data))

// }
