package main

//func main() {
//	n := make([]int, 3, 5)
//	fmt.Println("len=%d, cap=%d value=%v\n", len(n), cap(n), n)
//	n = append(n, 0, 0)
//	fmt.Println("len=%d, cap=%d value=%v\n", len(n), cap(n), n)
//	n = append(n, 1, 2, 3, 4, 5)
//	fmt.Println("len=%d, cap=%d value=%v\n", len(n), cap(n), n)
//
//	m := make([]int, 3)
//	fmt.Println("len=%d, cap=%d value=%v\n", len(m), cap(m), m)
//
//	b := make([]int, 0) // 0が入っている
//	var c []int         // メモリー確保しない、nil
//	fmt.Println("len=%d, cap=%d value=%v\n", len(b), cap(b), b)
//	fmt.Println("len=%d, cap=%d value=%v\n", len(c), cap(c), c)
//
//	//c = make([]int, 5)
//	c = make([]int, 0, 5)
//	for i := 0; i < 5; i++ {
//		c = append(c, i)
//		fmt.Println(c)
//	}
//	fmt.Println(c)
//}
