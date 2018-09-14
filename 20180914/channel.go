/*
在select中所有的case都没有匹配，则会走default分支
如果所有的case分支都没有满足条件，那么默认分支就会被选中并执行

如果没有默认分支，那么一旦所有的case 表达式都没有满足条件的，那么select语句就会被阻塞，直到至少有一个case表达式满足条件为止
当from循环和select嵌套使用的时候需要注意：
简单地在select语句中使用break语句，智能结束当前的select语句的执行，而并不会对外层的for语句产生作用，这种错误的用法会让这个for语句
无休止的一直运行下去
*/

package main

import (
	"fmt"
	"math/rand"
)

func main(){
	intChannels := [3] chan int {
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)
	fmt.Printf("the index :%d\n", index)

	intChannels[index] <- index
	select {
	case <-intChannels[0]:
		fmt.Println("the first candidate case is selected")
	case <- intChannels[1]:
		fmt.Println("the second candidate case is selected")
	case <- intChannels[2]:
		fmt.Println("the third candidate case is selected")
	default:
		fmt.Println("no candidate case is selected")
	}

}