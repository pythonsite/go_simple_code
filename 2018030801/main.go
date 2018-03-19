package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	if len(strs) == 1{
		return strs[0]
	}
	var minlen  = len(strs[0])
	var key int
	var minStr string
	for k,v := range strs{
		if len(v) <= minlen{
			minlen = len(v)
			key = k
			minStr = v
		}
	}
	strs = append(strs[:key],strs[key+1:]...)
	var prefixStr string
	BREAKPOINT:
		for k,v:=range []rune(minStr){
			for _,value:=range strs{
				if string(v) != string([]rune(value)[k]){
					break BREAKPOINT
				}
			}
			prefixStr += string(v)
		}
		return prefixStr
}

func main() {
	res := longestCommonPrefix([]string{"abc","abcd","abc","abc"})
	fmt.Println(res)
}
