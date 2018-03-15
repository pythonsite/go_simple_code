/*
平方差
 */

package main

import (
	"math"
	"fmt"
)

func getSum()float64{
	var num int
	for i:=1;i<=10;i++{
		num += i
	}
	sum := math.Pow(float64(num),2)
	return sum

}

func getSum2()float64{
	var num float64
	for i:=1;i<=10;i++{
		num += math.Pow(float64(i),2)
	}
	return num
}

func main() {
	num1 := getSum()
	fmt.Println(num1)
	num2 :=getSum2()
	fmt.Println(num2)
	res := num1-num2
	fmt.Println(res)
}
