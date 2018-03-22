package main

import (
	"strconv"
	"fmt"
	"math"
)

func reverse(x int) int{
	var b bool
	if x < 0{
		x = int(math.Abs(float64(x)))
		b = true
	}
	xStr := strconv.Itoa(x)
	xSlice := []rune(xStr)
	for i,j:=0,len(xSlice)-1;i<j;i,j = i+1,j-1{
		xSlice[i],xSlice[j] = xSlice[j],xSlice[i]
	}
	xresStr := string(xSlice)
	res,_ := strconv.Atoi(xresStr)
	if b == true{
		if res > math.MaxInt32 {
			return 0
		}
		return -res
	}
	if res > math.MaxInt32{
		return 0
	}else{
		return res
	}
}

func main() {
	res := reverse(-10)
	fmt.Println(res)
}