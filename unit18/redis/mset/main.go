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
	_, err = conn.Do("MSet", "name", "sss", "age", 17)
	if err != nil {
		fmt.Println("mset err...", err)
		return
	}
	fmt.Println("mset success...")

	// 3、从 Redis 读取数据
	r, err := redis.Strings(conn.Do("MGet", "name", "age"))
	if err != nil {
		fmt.Println("mget err...", err)
		return
	}
	// r -> interface{}
	fmt.Println("mget success...")
	fmt.Println("mget name age...", r) // mget name age... [sss 17]
}