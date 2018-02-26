/*
13195的因素数是5,7,13,29
求某个整数最大素数是多少
 */

package main

import "fmt"

//获取一个整数除了1所有的因数的切片
func getFactor(num int)(s []int) {
	s = append(s,num)
	for i := 2; i < num/2+1; i++ {
		if num%i == 0 {
			res := judge(i,s)
			if res == false{
				s = append(s,i,num/i)
			}

		}
	}
	return
}

// 判断n是否已经在s切片中
func judge(n int,s []int)(b bool){
	for _,v := range s{
		if v == n{
			return true
		}
	}
	return false
}

// 判断n是否是质数
func judgePrimeNum(s []int)(res []int){

	for _,v := range s{
		var b = true
		for i:=2;i<=v/2;i++{
			if v%i==0{
				b=false
				break
			}
		}
		if b == true{
			res = append(res, v)
		}
	}
	return
}

func getBig(s []int)(n int){
	bigNum := s[0]
	for i:=1;i<len(s);i++{
		if bigNum<s[i]{
			bigNum = s[i]
		}
	}
	return bigNum
}


func main() {
	s := getFactor(13195)
	fmt.Println(s)
	res := judgePrimeNum(s)
	fmt.Println(res)
	bigNum := getBig(res)
	fmt.Println(bigNum)

}
