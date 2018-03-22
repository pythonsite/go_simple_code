package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var res = []int{}
	for k,v := range nums{
		for j,q := range nums[k+1:]{
			if v + q == target{
				res = append(res,k,j+k+1)
				break
			}
		}
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums,26)
	fmt.Println(res)
}