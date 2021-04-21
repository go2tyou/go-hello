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
	_, err = conn.Do("HSet", "user1", "name", "ssr")
	if err != nil {
		fmt.Println("hset err...", err)
		return
	}
	_, err = conn.Do("HSet", "user1", "age", 17)
	if err != nil {
		fmt.Println("hset err...", err)
		return
	}
	fmt.Println("hset success...")

	// 3、从 Redis 读取数据
	r1, err := redis.String(conn.Do("HGet", "user1", "name"))
	if err != nil {
		fmt.Println("hget err...", err)
		return
	}
	r2, err := redis.Int(conn.Do("HGet", "user1", "age"))
	if err != nil {
		fmt.Println("hget err...", err)
		return
	}
	// r -> interface{}
	fmt.Println("hget success...")
	fmt.Println("hget name...", r1) // hget name... ssr
	fmt.Println("hget age...", r2) // hget age... 17
}