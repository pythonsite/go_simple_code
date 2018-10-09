package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp","192.168.0.116:6379")
	if err != nil {
		fmt.Println("conn redis failed:",err)
		return
	}
	fmt.Println("conn redis success")
	defer c.Close()
	_, err = c.Do("set","abc",100)
	if err != nil {
		fmt.Println("set error:",err)
		return
	}
	fmt.Println("set success")
	r, err := redis.Int(c.Do("get","abc"))
	if err != nil {
		fmt.Println("get error:",err)
		return
	}
	fmt.Println(r)
}