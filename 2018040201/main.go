/*
给定一个数组和一个值，在这个数组中原地移除指定值和返回移除后新的数组长度。

不要为其他数组分配额外空间，你必须使用 O(1) 的额外内存原地修改这个输入数组。

元素的顺序可以改变。超过返回的新的数组长度以外的数据无论是什么都没关系。

给定 nums = [3,2,2,3]，val = 3，

你的函数应该返回 长度 = 2，数组的前两个元素是 2。
 */

package  main

import "fmt"

//func removeElement(nums []int, val int) int {
//
//}

func main() {
	nums := [20]int{3,2,2,3}
	for k,v := range nums{
		if v == 3{
			nums[k] = 0
		}
	}
	fmt.Println(nums)
	fmt.Println(cap(nums))
}

