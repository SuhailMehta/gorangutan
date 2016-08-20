package main

import (
	"github.com/garyburd/redigo/redis"
)

type DbController struct {
	conn *redis.Conn
}

func NewRedisClient() *redis.Conn {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}

	return &client
}
