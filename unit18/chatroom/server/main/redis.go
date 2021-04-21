package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// 定义一个全局的连接池
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool {
		MaxIdle : maxIdle,
		MaxActive : maxActive,
		IdleTimeout : idleTimeout,
		Dial : func () (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}