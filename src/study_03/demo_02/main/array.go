package main

import "fmt"

func main() {
	
	sum := arr()
	fmt.Println("sum", sum)

	printArrInfo()

	var arr [3]int = [3]int{1, 2, 3}
	editArr(&arr)
	fmt.Println("edited arr", arr)
}


func arr() float64 {
	var hens [6]float64
	hens[0] = 1.0
	hens[1] = 1.0
	hens[2] = 1.0
	hens[3] = 1.0
	sum := 0.0
	for i := 0; i < 6; i++ {
		sum += hens[i]	
	}
	return sum
}

func printArrInfo() {
	var arr [3]int
	fmt.Println("arr", arr)
	// 数组的首地址与第一个元素的首地址一致
	fmt.Printf("arr 的首地址: %p, 第一个元素的地址: %p\n", &arr, &arr[0])
	
	// 指定下标
	var arr1 [3]int = [3]int{0: 1, 2: 2, 1: 3}
	fmt.Println("arr1", arr1)

    // for-range 遍历
	for index, value := range arr1 {
		fmt.Println("index: ", index, "value: ", value)
	}
}

// 形参必须写清除长度，不写的话就是切片类型
func editArr(arr *[3]int) {
	// 通过引用传递就可以直接修改原数组
	arr[0] = 100
}