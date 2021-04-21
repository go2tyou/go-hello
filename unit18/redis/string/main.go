package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	// 通过 go 向 redis 读写数据
	// 1、连接到 redis
	conn, err := redis.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("redis.Dial err...", err)
		return
	}

	defer conn.Close() // 关闭

	// fmt.Println("conn success...", conn)
	// 2、向 Redis 写入数据
	_, err = conn.Do("Set", "name", "ssr")
	if err != nil {
		fmt.Println("set err...", err)
		return
	}
	fmt.Println("set success...")

	// 3、从 Redis 读取数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err...", err)
		return
	}
	// r -> interface{}
	fmt.Println("get success...")
	// name := r.(string) // panic: interface conversion: interface {} is []uint8, not string
	fmt.Println("get name...", r) // get name... ssr
}