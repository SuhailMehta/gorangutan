package main

import (
	// "github.com/garyburd/redigo/redis"
	"log"
	"net/http"
)

func main() {

	redisClient := NewRedisClient()

	appC := DbController{conn: redisClient}

	c := *appC.conn

	defer c.Close()

	// c := &appC.conn

	// client.Do("SET", "best_car_ever", "Tesla Model S")

	router := appC.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
