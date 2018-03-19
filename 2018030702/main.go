package main

import "fmt"

func twoSum(nums []int, target int)[]int{
	if len(nums) <2 {
		return nil
	}
	m := make(map[int]int,len(nums))
	for i, v:= range nums{
		if j, ok := m[v];ok{
			return []int{j,i}
		}else{
			m[target-v] = i
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums,26)
	fmt.Println(res)
}