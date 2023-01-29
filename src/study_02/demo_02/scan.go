package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("请输入名字")
	fmt.Scanln(&name)
	fmt.Println("名字为", name)
}
