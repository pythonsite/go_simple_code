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
	defer c.Close()

	fmt.Println("con redis success")
	_, err = c.Do("hset","book","abc",100)
	if err != nil {
		fmt.Println("hset err:",err)
		return
	}
	r, err := redis.Int(c.Do("hget","book","abc"))
	if err != nil {
		fmt.Println("hget is err:",err)
		return
	}
	fmt.Println(r)
}