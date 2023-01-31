package main

import "fmt"

func main() {
	// 定义函数时就直接使用
	sum1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("sum1", sum1)

	// 使用变量接收匿名函数
	getSum := func (n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println("sum2", getSum(20, 20))
}