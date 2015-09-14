package main

import (
	"encoding/json"
	// "fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

		reorderedFollowersNotPtr := *reorderedFollowers

		GenerateCollage(reorderedFollowersNotPtr)

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

			reorderedFollowersNotPtr := *reorderedFollowers

			GenerateCollage(reorderedFollowersNotPtr)

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
		} else if ((value.Width + follower.Width) > defaultProporsions.Width) && ((value.MaxHeight + follower.Height) <= value.Height) {

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

type History struct {
	Height float64
	Width  float64
}

func GenerateCollage(followers followersLineType) {
	proportions := GetCanvasParameters(followers)

	m := image.NewRGBA(image.Rect(0, 0, int(proportions.Width), int(proportions.Height)))
	white := color.RGBA{255, 255, 255, 0}
	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.Point{0, 0}, draw.Src)

	heightIndent, widthIndent := 0.0, 0.0

	for _, line := range followers {
		maxHeight := 0.0
		lineHeight := 0.0

		smallHistory := []*History{}

		for _, follower := range line.Followers {

			response, err := http.Get(follower.ProfileImageUrl)

			if err != nil {
				log.Fatal("Image not loaded")
			}

			defer response.Body.Close()

			img, _, err := image.Decode(response.Body)

			if err != nil {
				log.Fatal(err)
			}

			newImage := resize.Resize(uint(follower.Width), uint(follower.Height), img, resize.Lanczos3)

			draw.Draw(m, m.Bounds(), newImage, image.Point{int(widthIndent) * -1, int(lineHeight+heightIndent) * -1}, draw.Src)

			widthIndent += follower.Width

			history := &History{Height: lineHeight + follower.Height, Width: widthIndent}

			smallHistory = append(smallHistory, history)

			if ((widthIndent + follower.Width) <= proportions.Width) && (follower.Height <= line.Height) {

				if maxHeight < (lineHeight + follower.Height) {
					maxHeight = lineHeight + follower.Height
				}

			} else if ((widthIndent + follower.Width) > proportions.Width) && ((maxHeight + follower.Height) <= line.Height) {

				lineHeight = maxHeight
				maxHeight += follower.Height
				widthIndent = line.StartWidth

				prevHeight := 0.0

				for i := (len(smallHistory) - 1); i >= 0; i-- {
					if smallHistory[i].Height <= follower.Height {
						prevHeight = smallHistory[i].Height

						smallHistory = smallHistory[:len(smallHistory)-1]
					} else {
						widthIndent = smallHistory[i].Width
						lineHeight = prevHeight
					}
				}

				log.Println("In 2 if")
			}
		}

		widthIndent = 0.0
		heightIndent = line.Height
	}

	out, err := os.Create("static/images/result.jpg")
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	log.Println("Image successful created")
}

func GetCanvasParameters(followers followersLineType) *Proportions {
	resultWidth, resultHeight := 0.0, 0.0

	for _, value := range followers {
		resultWidth += value.Width
		resultHeight += value.Height
	}

	return &Proportions{Width: resultWidth, Height: resultHeight}
}
