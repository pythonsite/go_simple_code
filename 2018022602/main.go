/*
回文数字，由两个两位位数字相乘获取的最大回文是9009=91*99
获取由两个三位数相乘获取的最大会回文数字
 */

package main

import (
	"fmt"
	"strconv"
)

// 将一个数字字符串进行反转
func reverseNum(num string)string{
	numSlice := []byte(num)
	for i,j:=0,len(numSlice)-1;i<j;i,j = i+1,j-1{
		numSlice[i],numSlice[j] = numSlice[j],numSlice[i]
	}
	res := string(numSlice)
	return res

}

// 判断反转后的字符串是否和原字符串相等
func judege(num string)(b bool){
	res := reverseNum(num)
	if num == res{
		return true
	}
	return false
}

func main() {
	var bigNum int
	var one int
	var two int
	for i:=100;i<=999;i++{
		for j:=100;j<=999;j++{
			num := i*j
			res := judege(strconv.Itoa(num))
			if res == true{
				if num > bigNum{
					bigNum = num
					one = i
					two = j
				}
			}
		}
	}
	fmt.Printf("%d*%d=%d",one,two,bigNum)

}
