package main

import (
	"bytes"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const GOOGLE_SEND_URL = "https://gcm-http.googleapis.com/gcm/send"
const API_KEY = "" // Add your api-key here

func RecoveryHandler(h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("recovery")
		h.ServeHTTP(rw, req)
	}

	return http.HandlerFunc(fn)
}

func LogHandler(h http.Handler) http.Handler {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, h)
}

/*
* POST request to send push notification to given deviceIds
* device_ids = 1,2,3,4
* title = "title to be shown for push"
* message = "message to be shown for push"
 */
func (client *DbController) AndroidPushNotification(rw http.ResponseWriter, req *http.Request) {
	dbSession := *client.conn

	deviceIds := req.FormValue("device_ids")
	title := req.FormValue("title")
	message := req.FormValue("message")

	if deviceIds == "" {
		rw.WriteHeader(http.StatusBadRequest)
	}

	//TODO: validation for proper comma separated structure and 1000 ids count, Duplicate device id check

	// Sending data to push server

	tokens := strings.Split(deviceIds, ",")

	var registrationIds string = ""

	for _, token := range tokens {

		tempValue, _ := redis.String(dbSession.Do("GET", token))

		registrationIds += (tempValue + ",")
	}

	registrationIds = strings.TrimRight(registrationIds, ",")
	rw.Write([]byte("[" + registrationIds + "]"))

	formattedString := `{
                          "data": {
                                   "title": "` + title + `",
                                   "message": "` + message + `"
                                   },
                           "registration_ids" : "` + registrationIds + `"
                         }`

	var jsonStr = []byte(formattedString)
	req, err := http.NewRequest("POST", GOOGLE_SEND_URL, bytes.NewBuffer(jsonStr))

	req.Header.Set("Authorization", "key="+API_KEY)
	req.Header.Set("Content-Type", "application/json")

	googleClient := &http.Client{}
	resp, err := googleClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	rw.Write([]byte(string(body)))

}

/*
 * Register device token to client specific registration token
 * Header = device-id - Unique device token
 */
func (client *DbController) RegisterDevice(rw http.ResponseWriter, req *http.Request) {
	dbSession := *client.conn

	header := req.Header.Get("device-id")

	if header == "" {
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	registrationId := req.FormValue("registration_id")

	if registrationId == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	dbSession.Do("SET", header, registrationId)

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(header))

}

// For future purpose
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
