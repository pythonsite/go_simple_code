/*
通过计算一定数量的斐波那契数列，并求出偶数和
 */

package main

import "fmt"

// 用于计算获取斐波那契数列
func fib(n int)int{
	if n<=2{
		return n
	}else{
		return fib(n-1)+fib(n-2)
	}
}

func main() {
	var s []int
	var sum int
	for i:=1;i<10;i++{
		res := fib(i)
		if res % 2 == 0{
			sum += res
		}
		fmt.Println(res)
		s = append(s,res)
	}
	fmt.Println(s)
	fmt.Println(sum)
}