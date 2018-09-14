/*
这个例子用于理解select的用法


*/

package main


import (
	"time"
	"fmt"
)

func main(){
	intChan := make(chan int, 1)
	
	//声明了一个定时器，1秒之后关闭intChan
	time.AfterFunc(2*time.Second, func(){
		close(intChan)
	})
	
	// 因为初始的intChan中并没有任何元素，所有case语句会阻塞，并且没有default
	// 直到我们的的定时器关闭chan后，我们唯一的case才满足条件执行
	select {
	case _, ok := <- intChan:
		if !ok {
			fmt.Println("the candidate case is closed")
			break
		}
		fmt.Println("the candidate case is selected")
	}
}