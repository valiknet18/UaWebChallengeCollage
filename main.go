package main

import (
	// "UaWebChallenge/twitterApi"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	tokens map[string]*oauth.RequestToken
	c      *oauth.Consumer
	client *http.Client
)

func main() {
	tokens = make(map[string]*oauth.RequestToken)

	config := parseConfig()

	c = oauth.NewConsumer(
		config.ConsumerApiKey,
		config.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)

	r := httprouter.New()

	r.GET("/", IndexAction)
	r.GET("/api/auth", RedirectUserToTwitter)
	r.GET("/api/maketoken", GetTwitterToken)
	r.GET("/api/followers", GetFollowersAction)
	r.GET("/api/verify_auth", VerifyAuthAction)

	r.ServeFiles("/static/*filepath", http.Dir("static/"))

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", r)
}

func IndexAction(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	page, err := ioutil.ReadFile("static/index.html")

	if err != nil {
		log.Fatal("Static page not loaded")
	}

	fmt.Fprintf(rw, string(page))
}

func VerifyAuthAction(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	type resultMessage struct {
		Status int `json:"status"`
	}

	if client != nil {
		data := &resultMessage{Status: 200}

		result, err := json.Marshal(data)

		if err != nil {
			log.Fatal("not parsed")
		}

		rw.Header().Set("Content-type", "application/json")
		rw.Write(result)
	} else {
		data := &resultMessage{Status: 401}

		result, err := json.Marshal(data)

		if err != nil {
			log.Fatal("not parsed")
		}

		rw.Header().Set("Content-type", "application/json")
		rw.Write(result)
	}
}

func GetFollowersAction(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	response, err := client.Get("https://api.twitter.com/1.1/followers/list.json?cursor=-1&screen_name=v1per16&skip_status=true&include_user_entities=false")

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(rw, "The newest item in your home timeline is: "+string(bits))
}

//This return auth user
func RedirectUserToTwitter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// values := r.URL.Query()

	// if values.Get("oauth_verifier") != "" && values.Get("oauth_token") != "" {

	// }

	tokenUrl := fmt.Sprintf("http://%s/api/maketoken", r.Host)
	token, requestUrl, err := c.GetRequestTokenAndUrl(tokenUrl)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure to save the token, we'll need it for AuthorizeToken()
	tokens[token.Token] = token

	http.Redirect(w, r, requestUrl, http.StatusTemporaryRedirect)
}

//This func make return Client
func GetTwitterToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	values := r.URL.Query()
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")

	accessToken, err := c.AuthorizeToken(tokens[tokenKey], verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	client, err = c.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "http://localhost:8000/#/", http.StatusTemporaryRedirect)
}

/**
- width = 400
- height = 300
- count = 50
*/
