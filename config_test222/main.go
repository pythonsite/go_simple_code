package main

import (
	"go_dev/config"
	"fmt"
)

func main() {
	c := config.ConfigEngine{}
	c.Load("test.yaml")
	fmt.Println(c.GetInt("Main.HttpPort"))
}

