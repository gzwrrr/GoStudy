package main

import (
	"fmt"
)

var (
	g1 = 10
	g2 = "heehh"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("ahahhaa")

	var i int = 10
	fmt.Println(i)

	var str = "hello"
	fmt.Println(str)

	str2 := "hello2"
	fmt.Println(str2)

	n1, n2, n3 := 10, "str", "str2"
	fmt.Println("n1", n1, "n2", n2, "n3", n3)

	var f1 float32 = 32.32
	var f2 float64 = .32
	var f3 float64 = .32e2
	fmt.Println("f1", f1, "f2", f2, "f3", f3)

	var c1 byte = 'a'
	fmt.Println("c1", c1)

	var b1 bool = false
	fmt.Println("b1", b1)

	var int1 int32 = 32
	var float1 float32 = float32(int1)
	fmt.Println("int1", int1, "float1", float1)
}
