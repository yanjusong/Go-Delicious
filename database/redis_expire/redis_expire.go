package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()

	// 设置abc的值在10秒钟后过期
	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}
