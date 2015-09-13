package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
)

var usersFollowers = make(map[string]*Followers)

type Followers struct {
	Followers []*Follower `json:"users"`
}

type Follower struct {
	Id              int    `json:"id"`
	IdStr           string `json:"id_str"`
	ProfileImageUrl string `json:"profile_image_url"`
	StatusesCount   int    `json:"statuses_count"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
}

//Return user followers
func GetFollowersAction(name string) *Followers {
	// if val, ok := usersFollowers[name]; ok {
	// 	return val
	// } else {
	// 	usersFollowers[name] = &Followers{}

	// 	return usersFollowers[name]
	// }

	if val, ok := usersFollowers[name]; ok {
		return val
	} else {
		if client != nil {
			response, err := client.Get("https://api.twitter.com/1.1/followers/list.json?cursor=-1&screen_name=" + name + "&skip_status=true&include_user_entities=false&count=100")

			if err != nil {
				log.Fatal(err)
			}

			defer response.Body.Close()

			bits, err := ioutil.ReadAll(response.Body)

			followers := &Followers{}

			json.Unmarshal(bits, followers)

			usersFollowers[name] = followers

			return followers
		} else {
			log.Fatal("client variable not initialized")

			return &Followers{}
		}
	}
}
