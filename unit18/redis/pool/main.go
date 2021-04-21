package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 定义一个全局的连接池
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init()  {
	pool = &redis.Pool {
		MaxIdle : 8,
		MaxActive : 0,
		IdleTimeout : 100,
		Dial : func () (redis.Conn, error) {
			return redis.Dial("tcp", "0.0.0.0:6379")
		},
	}
}

func main()  {
	// 从连接池取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("set err...", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err...", err)
		return
	}
	fmt.Println("get suc...", r) // get suc... tom

	// 连接池关闭
	pool.Close()
	// conn2 := pool.Get()
	// defer conn2.Close()
	// fmt.Println("conn2...", conn2)
	// _, err = conn2.Do("Set", "name", "tom")
	// if err != nil {
	// 	fmt.Println("set err...", err) // set err... redigo: get on closed pool
	// 	return
	// }
}