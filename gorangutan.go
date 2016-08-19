package main

import (
	// "github.com/justinas/alice"
	"log"
	"net/http"
)

func main() {

	// commonHandler := alice.New(loggingHandler, recoveryHandler)

	// myHandler := http.HandlerFunc(androidGCM)
	// http.Handle("/android", commonHandler.Then(myHandler))

	// http.ListenAndServe(":8080", nil)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8088", router))

}
