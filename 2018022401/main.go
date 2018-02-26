/*
求出1000以内既能被3整除也能被5整除的所有数的和
*/
package main

import "fmt"

func main(){
	var num int
	for i:=1;i<=1000;i++{
		if i%3==0 || i%5==0{
			num += i
		}
	}
	fmt.Println(num)
}
