package main

import (
	"fmt"
)

func main() {
	var i int = 10
	fmt.Println("i", i)
	fmt.Println("i 的地址", &i)

	var p *int = &i
	fmt.Println("p 的地址", &p)
	fmt.Println("p 指向的地址", p)
	// 解引用
	fmt.Println("p 的值", *p)

}
