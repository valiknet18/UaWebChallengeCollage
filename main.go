package main

import (
	"UaWebChallenge/twitterApi"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	config := parseConfig()

	twitterApi.Init(config.ConsumerApiKey, config.ConsumerSecret)

	r := httprouter.New()

	r.ServeFiles("/*filepath", http.Dir("static/"))

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", r)
}

/**
- width = 400
- height = 300
- count = 50
*/
