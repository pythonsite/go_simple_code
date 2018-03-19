package main

import "fmt"

func removeElement(nums []int, val int) int {
	var sum int
	for _,v:=range nums{
		if val == v{
			sum += 1
		}
	}
	var count int
	for {
		for k,v:=range nums{
			if val == v{
				count += 1
				nums = append(nums[:k],nums[k+1:]...)
				break
			}
		}
		if count == sum{
			break
		}

	}
	return len(nums)
}

func main() {
	nums := []int{3,2,4,3}
	res := removeElement(nums, 3)
	fmt.Println(res)
}
