package main

import "fmt"

func reverse(runes []rune) []rune{
	// 将切片内容进行翻转
	for i,j:=0,len(runes)-1;i<j;i,j = i+1,j-1{
		runes[i],runes[j] = runes[j],runes[i]
	}
	return runes
}

func convertToSlice(str string) []rune{
	// 将字符串转换为slice
	res := []rune(str)
	return res
}

func main() {
	str := "-123"
	res := convertToSlice(str)
	fmt.Println(string(res))
	res2 := reverse(res)
	fmt.Println(string(res2))
}
