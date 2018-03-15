/*
2520是可以被1到10中的每个数字除以没有任何余数的最小数字
可以被1到20的所有数字整除的最小正数是多少
 */
package main

import "fmt"

func getSlice(num int) ([]int) {
	var s = []int{}
	var s2 = []int{}
	for i := 1; i <= num; i++ {
		s = append(s, i)
	}
	for j := 2; j < num; j++ {
		flag := 1
		for {
			flag = 0
			for k := 0; k < 20; k++ {
				if s[k]%j == 0 {
					s[k] = s[k] / j
					flag = 1
				}
			}
			if flag == 1 {
				s2 = append(s2, j)
			} else {
				break
			}
		}

	}
	return s2
}

func main() {
	s := getSlice(20)
	var result = 1
	for _,v:=range s{
		result*=v
	}
	fmt.Println(result)
}
