/*
什么是高阶函数
1、接收其他的函数作为参数传入
2、把其他的函数作为结果返回
*/


package main

import (
	"fmt"
	"errors"
)


type Printer func(content string) (n int, err error)


func printToStd(content string) (n int, err error) {
	return fmt.Println(content)
}

type operate func(x, y int) int


func calculate(x int ,y int, op operate) (int, error){
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}


func main() {
	var p Printer
	p = printToStd
	p("something")


	op := func(x,y int) int {
		return x+y
	}
	res,_ := calculate(10,20,op)
	fmt.Println(res)
}