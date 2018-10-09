/*
关于巴贝奇差分机
*/

package main

import (
	"fmt"
)

func bafenqi(n int)int {
	square_nminus1 := 1
	square_n := 4
	alpha_nminus1 := 1
	
	for i:=2;i<=n-1;i++ {
		alpha_n := square_n - square_nminus1
		beta_n := alpha_n - alpha_nminus1
		square_nplus1 := square_n + alpha_n + beta_n
		square_nminus1 = square_n
		square_n = square_nplus1
		alpha_nminus1 = alpha_n
	}
	return square_n
}

func main(){
	res := bafenqi(10)
	fmt.Println(res)
}