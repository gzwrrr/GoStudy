package main


import (
	"fmt"
	"study_02/demo_03/utils"
)

func init() {
	fmt.Println("init....")
}

func getSum(num1 float64, num2 float64) float64 {
	return num1 + num2
}

type myGetSumType func(float64, float64) float64
func myFunc(myGetSum  myGetSumType, args ...float64) (sum float64) {
	sum = myGetSum(args[0], args[1])	
	return 
}

func main() {
	var sum, flag = utils.Add(1.1, 2.2)
	fmt.Println("sum", sum, flag)
	fmt.Println("sum", myFunc(getSum, 2.2, 3.3))
}

