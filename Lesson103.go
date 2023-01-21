package main

const Pi = 3.14

// 他のファイルから呼び出す大文字で書く

const (
	Username = "test_user"
	Password = "test_pass"
)

// var big int = 9223372036854775807 + 1 // 型をした場合、overflowとなる。
//var big = 9223372036854775807 + 1 // 型を指定しなければ、エラーとならない

// int64 環境指定

//func main() {
//	fmt.Println(Pi, Username, Password)
//	// Pi = 3 // エラーとなる、書き換えられない
//	fmt.Println(big - 1) //宣言時にエラーとなる
//}
