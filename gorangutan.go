package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	redisClient := NewRedisClient()

	appC := DbController{conn: redisClient}

	c := *appC.conn

	defer c.Close()

	router := appC.NewRouter()

	serve := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        LogHandler(router),
	}

	log.Fatal(serve.ListenAndServe())

}
