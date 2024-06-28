package main

import (
	"math"
)

type ColorRange struct {
	Min, Max [3]int
}

var comprehensiveColorRanges = map[string]ColorRange{
	"black":      {Min: [3]int{0, 0, 0}, Max: [3]int{64, 64, 64}},
	"dark_gray":  {Min: [3]int{65, 65, 65}, Max: [3]int{128, 128, 128}},
	"light_gray": {Min: [3]int{129, 129, 129}, Max: [3]int{191, 191, 191}},
	"white":      {Min: [3]int{192, 192, 192}, Max: [3]int{255, 255, 255}},

	"dark_red":  {Min: [3]int{129, 0, 0}, Max: [3]int{191, 64, 64}},
	"red":       {Min: [3]int{192, 0, 0}, Max: [3]int{255, 64, 64}},
	"light_red": {Min: [3]int{255, 65, 65}, Max: [3]int{255, 128, 128}},
	"pink":      {Min: [3]int{192, 0, 129}, Max: [3]int{255, 127, 191}},

	"dark_green":  {Min: [3]int{0, 129, 0}, Max: [3]int{64, 191, 64}},
	"green":       {Min: [3]int{0, 192, 0}, Max: [3]int{64, 255, 64}},
	"light_green": {Min: [3]int{65, 192, 65}, Max: [3]int{191, 255, 191}},

	"dark_blue":  {Min: [3]int{0, 0, 129}, Max: [3]int{64, 64, 191}},
	"blue":       {Min: [3]int{0, 0, 192}, Max: [3]int{64, 64, 255}},
	"light_blue": {Min: [3]int{65, 65, 192}, Max: [3]int{191, 191, 255}},

	"dark_yellow":  {Min: [3]int{129, 129, 0}, Max: [3]int{191, 191, 64}},
	"yellow":       {Min: [3]int{192, 192, 0}, Max: [3]int{255, 255, 64}},
	"light_yellow": {Min: [3]int{192, 192, 65}, Max: [3]int{255, 255, 191}},

	"dark_cyan":  {Min: [3]int{0, 129, 129}, Max: [3]int{64, 191, 191}},
	"cyan":       {Min: [3]int{0, 192, 192}, Max: [3]int{64, 255, 255}},
	"light_cyan": {Min: [3]int{65, 192, 192}, Max: [3]int{191, 255, 255}},

	"dark_magenta":  {Min: [3]int{129, 0, 129}, Max: [3]int{191, 64, 191}},
	"magenta":       {Min: [3]int{192, 0, 192}, Max: [3]int{255, 64, 255}},
	"light_magenta": {Min: [3]int{192, 65, 192}, Max: [3]int{255, 191, 255}},

	"dark_orange":  {Min: [3]int{129, 65, 0}, Max: [3]int{191, 128, 64}},
	"orange":       {Min: [3]int{192, 65, 0}, Max: [3]int{255, 128, 64}},
	"light_orange": {Min: [3]int{192, 129, 65}, Max: [3]int{255, 191, 128}},

	"dark_purple":  {Min: [3]int{65, 0, 129}, Max: [3]int{128, 64, 191}},
	"purple":       {Min: [3]int{129, 0, 192}, Max: [3]int{191, 64, 255}},
	"light_purple": {Min: [3]int{129, 65, 192}, Max: [3]int{255, 191, 255}},

	"dark_brown":  {Min: [3]int{65, 32, 0}, Max: [3]int{128, 96, 64}},
	"brown":       {Min: [3]int{129, 32, 0}, Max: [3]int{191, 96, 64}},
	"light_brown": {Min: [3]int{129, 97, 65}, Max: [3]int{255, 191, 128}},
}

func classifyColor(rgbPoint [3]float64, colorRanges map[string]ColorRange) string {
	minDistance := math.Inf(1)
	nearestColor := "unknown"

	for color, ranges := range colorRanges {
		center := [3]float64{
			float64(ranges.Min[0]+ranges.Max[0]) / 2,
			float64(ranges.Min[1]+ranges.Max[1]) / 2,
			float64(ranges.Min[2]+ranges.Max[2]) / 2,
		}

		distance := math.Sqrt(
			math.Pow(rgbPoint[0]-center[0], 2) +
				math.Pow(rgbPoint[1]-center[1], 2) +
				math.Pow(rgbPoint[2]-center[2], 2),
		)

		if distance < minDistance {
			minDistance = distance
			nearestColor = color
		}
	}

	return nearestColor
}

//func main() {
//	// Example usage
//	rgbPoint := [3]float64{255, 0, 0}
//	colorName := classifyColor(rgbPoint, fiftyColorRanges)
//	fmt.Printf("The color RGB(%v, %v, %v) is classified as: %s\n", rgbPoint[0], rgbPoint[1], rgbPoint[2], colorName)
//}
