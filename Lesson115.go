package main

import "fmt"

func add(x int, y int) int {
	// func add(x, y int) int { ともかける
	fmt.Println("add function")
	fmt.Println(x + y)
	return x + y
}

// 2つ返り値する
func add2(x int, y int) (int, int) {
	// func add(x, y int) int { ともかける
	fmt.Println("add function")
	fmt.Println(x + y)
	return x + y, x - y
}

func cal(price, item int) (result int) {
	// result := price * item // 返り値で定義されたので、宣言できない
	result = price * item
	return result // reusltを省略できる
}

// 名付けることによって、可動性上がる

func conver(price int) float64 {
	return float64(price)
}

// 読みやすい場合は、不要と考えられる
//
//func main() {
//	r := add(10, 20)
//	fmt.Println(r)
//
//	r1, r2 := add2(10, 20)
//	fmt.Println(r1, r2)
//
//	r3 := cal(100, 2)
//	fmt.Println(r3)
//
//	f := func() {
//		fmt.Println("inner func")
//	}
//	f()
//
//	g := func(x int) {
//		fmt.Println("inner func", x)
//	}
//	g(1)
//
//	func(x int) {
//		fmt.Println("inner func", x)
//	}(3)
//
//	// go func で並列化するときに利用する
//}
