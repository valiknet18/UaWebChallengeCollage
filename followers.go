package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"sort"
)

var usersFollowers = make(map[string]*Followers)

type Followers struct {
	Followers []*Follower `json:"users"`
}

type ByStatusesCount []*Follower

func (a ByStatusesCount) Len() int           { return len(a) }
func (a ByStatusesCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStatusesCount) Less(i, j int) bool { return a[i].StatusesCount > a[j].StatusesCount }

type Follower struct {
	Id              int     `json:"id"`
	IdStr           string  `json:"id_str"`
	ProfileImageUrl string  `json:"profile_image_url"`
	StatusesCount   int     `json:"statuses_count"`
	Name            string  `json:"name"`
	ScreenName      string  `json:"screen_name"`
	Width           float64 `json:"width"`
	Height          float64 `json:"height"`
	NewLine         bool    `json:"new_line"`
}

type FollowersLine struct {
	StartWidth float64     `json:"start_width"`
	Width      float64     `json:"width"`
	Height     float64     `json:"heigh"`
	LineHeight float64     `json:"line_height"`
	MaxHeight  float64     `json:"max_height"`
	Followers  []*Follower `json:"followers"`
}

type followersLineType []*FollowersLine

//Return user followers
func GetFollowersAction(name string, proportions Proportions) *followersLineType {
	if val, ok := usersFollowers[name]; ok {
		var maxStatusCount = 0

		for _, value := range val.Followers {
			if value.StatusesCount > maxStatusCount {
				maxStatusCount = value.StatusesCount
			}
		}

		for _, value := range val.Followers {
			persentProportialsStatuses := (float64(value.StatusesCount) / float64(maxStatusCount)) * 100.0

			getWidthAndHeighPhoto(persentProportialsStatuses, proportions, value)
		}

		followersNotPtr := *val

		reorderedFollowers := ReorderHash(proportions, followersNotPtr)

		return reorderedFollowers
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

			var maxStatusCount = 0

			for _, value := range followers.Followers {
				if value.StatusesCount > maxStatusCount {
					maxStatusCount = value.StatusesCount
				}
			}

			for _, value := range followers.Followers {
				persentProportialsStatuses := (float64(value.StatusesCount) / float64(maxStatusCount)) * 100.0

				getWidthAndHeighPhoto(persentProportialsStatuses, proportions, value)

			}

			sort.Sort(ByStatusesCount(followers.Followers))

			usersFollowers[name] = followers

			followersNotPtr := *followers

			reorderedFollowers := ReorderHash(proportions, followersNotPtr)

			return reorderedFollowers
		} else {
			log.Fatal("client variable not initialized")

			return &followersLineType{}
		}
	}
}

func ReorderHash(defaultProporsions Proportions, followers Followers) *followersLineType {
	var resultFollowers = followersLineType{}

	for _, value := range followers.Followers {
		resultFollowers = FindAndInsertInSlice(value, defaultProporsions, resultFollowers)
	}

	// fmt.Println(resultFollowers)

	return &resultFollowers
}

func FindAndInsertInSlice(follower *Follower, defaultProporsions Proportions, followers followersLineType) followersLineType {
	flag := false

	for _, value := range followers {
		if ((value.Width + follower.Width) <= defaultProporsions.Width) && (follower.Height <= value.Height) {
			value.Followers = append(value.Followers, follower)

			value.Width += follower.Width

			if value.LineHeight+follower.Height > value.MaxHeight {
				value.MaxHeight = value.LineHeight + follower.Height
			}

			flag = true

			break
		} else if ((value.Width + follower.Width) > defaultProporsions.Width) && ((value.MaxHeight + follower.Height) < value.Height) {

			follower.NewLine = true

			value.Followers = append(value.Followers, follower)
			value.Width = value.StartWidth + follower.Width

			value.LineHeight = value.MaxHeight
			value.MaxHeight += follower.Height

			flag = true

			break
		}
	}

	if !flag {
		newLine := &FollowersLine{LineHeight: 0, StartWidth: follower.Width, Width: follower.Width, Height: follower.Height, MaxHeight: 0}
		newLine.Followers = append(newLine.Followers, follower)

		followers = append(followers, newLine)
	}

	return followers
}
