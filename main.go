package main

import (
	"UaWebChallenge/twitterApi"
	"fmt"
	// "github.com/julienschmidt/httprouter"
	// "log"
	// "net/http"
)

func main() {
	config := parseConfig()

	twiteerApi.Init(config.ConsumerApiKey, config.ConsumerSecret)

	// _ := httprouter.New()

	// r.ServeFiles("/*filepath", http.Dir("static/"))

	// fmt.Println("Server running on port :8000")
	// http.ListenAndServe(":8000", r)
}

/**
- width = 400
- height = 300
- count = 50
*/
