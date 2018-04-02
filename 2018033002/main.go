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
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}

func main() {
	res := reverse(123)
	fmt.Println(res)
}
