package main

import "fmt"

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println("n", f(1)) // 11
	fmt.Println("n", f(1)) // 12
	fmt.Println("add", AddUpper()(1)) // 11
	fmt.Println("add", AddUpper()(1)) // 11
}