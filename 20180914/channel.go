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