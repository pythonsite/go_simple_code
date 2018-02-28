package main

import (
	"fmt"
)

func main() {
	var str = "789345"
	for i:=0;i<=len(str)-2;i++{
		num := str[i:i+2]
		fmt.Println(num)
	}
}