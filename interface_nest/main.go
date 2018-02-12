package main

import "fmt"

// 这里定义一个Eater接口
type Eater interface {
	Eat()
}

// 这里定义一个Talker接口
type Talker interface {
	Talk()
}

// 这里定义个动物的接口，同时嵌套了Eater和Talker接口
type Animal interface {
	Eater
	Talker
}

// 这里定义一个Dog的struct，并实现talk方法和eat方法，这样就实现了动物的接口
type Dog struct {

}

func (d *Dog) Talk(){
	fmt.Println("talk....")
}

func (d *Dog) Eat(){
	fmt.Println("eating....")
}

func main() {
	d := &Dog{}
	var a Animal
	a = d
	a.Eat()
	a.Talk()
}