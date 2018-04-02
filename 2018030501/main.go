/*
给定一个范围为 32 位 int 的整数，将其颠倒。
如
输入: 123
输出:  321

输入: -123
输出: -321

输入: 120
输出: 21

假设我们的环境只能处理 32 位 int 范围内的整数。根据这个假设，如果颠倒后的结果超过这个范围，则返回 0。
 */


package main

import (
	"strconv"
	"fmt"
	"math"
)

func reverse(x int) int{
	var b bool
	if x < 0{
		x = int(math.Abs(float64(x)))
		b = true
	}
	xStr := strconv.Itoa(x)
	xSlice := []rune(xStr)
	for i,j:=0,len(xSlice)-1;i<j;i,j = i+1,j-1{
		xSlice[i],xSlice[j] = xSlice[j],xSlice[i]
	}
	xresStr := string(xSlice)
	res,_ := strconv.Atoi(xresStr)
	if b == true{
		if res > math.MaxInt32 {
			return 0
		}
		return -res
	}
	if res > math.MaxInt32{
		return 0
	}
	return res
	
}

func main() {
	res := reverse(-10)
	fmt.Println(res)
}