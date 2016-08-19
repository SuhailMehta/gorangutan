package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func RecoveryHandler(h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("recovery")
		h.ServeHTTP(rw, req)
	}
	//timeEntry := time.Now()

	return http.HandlerFunc(fn)
}

func LoggingHandler(h http.Handler) http.Handler {
	// fn := func(rw http.ResponseWriter, req *http.Request) {
	// 	h.ServeHTTP(rw, req)
	// }
	//timeEntry := time.Now()
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, h)
}

func AndroidGCM(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("androidGCM")
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
