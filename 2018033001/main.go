/*

给定一个字符串， 包含大小写字母、空格 ' '，请返回其最后一个单词的长度。

如果不存在最后一个单词，请返回 0 。

注意事项：一个单词的界定是，由字母组成，但不包含任何的空格。
 */

package main

import (
	"strings"
	"fmt"
)

func lengthOfLastWord(s string) int {
	s  = strings.TrimSpace(s)
	arr := strings.Split(s," ")
	res := arr[len(arr)-1]
	return len(res)
}

func main() {
	res := lengthOfLastWord("hello world")
	fmt.Println(res)
}