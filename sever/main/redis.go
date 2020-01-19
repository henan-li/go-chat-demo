package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, idleTimeOut time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeOut,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", address)
		},
	}
}
