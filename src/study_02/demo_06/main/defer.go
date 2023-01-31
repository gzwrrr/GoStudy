package main

import "fmt"


func f() {
	// 当执行到 defer 时，暂时不会执行，会压入到一个独立的栈中
	// 当函数执行完毕后，defer 标识的操作会以先入后出的顺序执行
	defer fmt.Println("结束操作 2")
	defer fmt.Println("关闭资源 1")
	fmt.Println("执行操作...")
}

func main() {
	f()
}