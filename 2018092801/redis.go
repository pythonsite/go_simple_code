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
}