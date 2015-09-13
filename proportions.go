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
	var coefficient = 1.0

	if persentProportialsStatuses > 80.0 {
		coefficient = 0.8
	} else if persentProportialsStatuses > 60.0 {
		coefficient = 0.6
	} else if persentProportialsStatuses > 40.0 {
		coefficient = 0.4
	} else {
		coefficient = 0.2
	}

	MaxWidthImage := (defaultProportions.Width / 100) * 60
	MaxHeightImage := (defaultProportions.Height / 100) * 60

	follower.Width = MaxWidthImage * coefficient
	follower.Height = MaxHeightImage * coefficient
}
