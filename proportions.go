package main

import (
// "fmt"
)

type Proportions struct {
	Width  float64
	Height float64
}

func setProportions(width, height float64) Proportions {
	proportions := Proportions{}
	proportions.Width = width
	proportions.Height = height

	return proportions
}

func getWidthAndHeighPhoto(persentProportialsStatuses float64, defaultProportions Proportions, follower *Follower) {
	//It's 100%

	var coefficient = 1.0

	if persentProportialsStatuses > 80.0 {
		coefficient = 1.03
	} else if persentProportialsStatuses > 60.0 {
		coefficient = 1.20
	} else if persentProportialsStatuses > 40.0 {
		coefficient = 1.40
	} else if persentProportialsStatuses > 20.0 {
		coefficient = 2.35
	} else if persentProportialsStatuses > 10.0 {
		coefficient = 4.6
	} else if persentProportialsStatuses > 5.0 {
		coefficient = 7.5
	} else if persentProportialsStatuses > 2.0 {
		coefficient = 22.20
	} else if persentProportialsStatuses > 1.0 {
		coefficient = 38
	} else if persentProportialsStatuses > 0.5 {
		coefficient = 100
	} else {
		coefficient = 150
	}

	MaxWidthImage := ((defaultProportions.Width - 200) / 2) * 1
	MaxHeightImage := ((defaultProportions.Height - 200) / 2) * 1

	follower.Width = (MaxWidthImage / 100) * persentProportialsStatuses * coefficient
	follower.Height = (MaxHeightImage / 100) * persentProportialsStatuses * coefficient
}
