package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp","192.168.0.116:6379")
	if err != nil {
		fmt.Println("con redis err:",err)
		return
	}
	fmt.Println("con redis success")
	defer c.Close()

	_, err = c.Do("lpush","book_list","abcd","age",23)
	if err != nil {
		fmt.Println("redis lpush err:",err)
		return
	}

	r, err := redis.String(c.Do("lpop","book_list"))
	if err !=nil {
		fmt.Println("redis lpop is err:",err)
		return
	}
	fmt.Println(r)

}