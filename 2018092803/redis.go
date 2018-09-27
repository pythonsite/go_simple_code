package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp","192.168.0.116:6379")
	if err != nil {
		fmt.Println("conn redis err:",err)
		return
	}
	fmt.Println("conn redis success")
	defer c.Close()

	_, err = c.Do("mset","abc",23,"age",1212)
	if err != nil {
		fmt.Println("mset error:",err)
		return
	}
	fmt.Println("mset success")

	r, err := redis.Ints(c.Do("mget","abc","age"))
	if err != nil {
		fmt.Println("redis mget err:",err)
		return
	}
	for _, v := range r{
		fmt.Println(v)
	}
}