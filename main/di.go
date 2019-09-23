package main

import (
	"fmt"
	"io"
)

// 把第一个参数从io.Writer的实现改为接口本身，可以把字符串写到不同的地方，增加了通用性
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
//	Greet(w, "world")
//}

//func main()  {
//	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
//}
