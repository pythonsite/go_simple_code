/*
回去第1001个素数
 */

package main

import "fmt"

func main() {
	var num int
	for i:=2;;i++{
		var b bool
		for j:=2;j<=i/2;j++{
			if i%j == 0{
				b = true
				break
			}
		}
		if b == false{
			num += 1
			if num == 1001{
				fmt.Println(i)
				break
			}
		}
	}
}

