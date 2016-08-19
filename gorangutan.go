package main

import (
	// "github.com/garyburd/redigo/redis"
	"log"
	"net/http"
)

func main() {

	client := NewRedisClient()

	defer client.Close()

	// appC := DbController{conn: &client}

	// c := &appC.conn

	client.Do("SET", "best_car_ever", "Tesla Model S")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8088", router))

}
