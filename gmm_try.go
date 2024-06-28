package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"sort"
)

import (
	_ "image/jpeg"
	_ "image/png"
)

type Color struct {
	R, G, B float64
}

func visualizeColors(colors []Color, imagePath string) {
	// This function would require a Go graphics library to implement
	// Consider using a library like "github.com/fogleman/gg" for drawing
	fmt.Println("Visualization not implemented in this Go version")
}

func findDominantColors(imagePath string, plot bool, weightScale float64, minDuplicates int) ([]Color, error) {
	// Open the image
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a slice to store all pixels
	var pixels []Color

	// Iterate through all pixels
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels = append(pixels, Color{
				R: float64(r >> 8),
				G: float64(g >> 8),
				B: float64(b >> 8),
			})
		}
	}

	// Define extreme colors
	extremeColors := []Color{
		{R: 255, G: 255, B: 255},
		{R: 0, G: 0, B: 0},
	}

	// Calculate weights
	weights := make([]float64, len(pixels))
	for i, pixel := range pixels {
		minDist := math.Inf(1)
		for _, ec := range extremeColors {
			dist := math.Sqrt(math.Pow(pixel.R-ec.R, 2) + math.Pow(pixel.G-ec.G, 2) + math.Pow(pixel.B-ec.B, 2))
			if dist < minDist {
				minDist = dist
			}
		}
		weights[i] = math.Exp(-minDist * weightScale)
	}

	// Normalize weights
	sum := 0.0
	for _, w := range weights {
		sum += w
	}
	for i := range weights {
		weights[i] /= sum
	}

	// Create weighted pixels
	var weightedPixels []Color
	for i, pixel := range pixels {
		duplicates := int(weights[i] * 100)
		if duplicates < minDuplicates {
			duplicates = minDuplicates
		}
		for j := 0; j < duplicates; j++ {
			weightedPixels = append(weightedPixels, pixel)
		}
	}

	// Here, you would implement a Gaussian Mixture Model or use an alternative clustering method
	// For simplicity, let's just return the 3 most common colors
	colorCounts := make(map[Color]int)
	for _, pixel := range weightedPixels {
		colorCounts[pixel]++
	}

	var sortedColors []Color
	for color := range colorCounts {
		sortedColors = append(sortedColors, color)
	}

	sort.Slice(sortedColors, func(i, j int) bool {
		return colorCounts[sortedColors[i]] > colorCounts[sortedColors[j]]
	})

	if len(sortedColors) > 3 {
		sortedColors = sortedColors[:3]
	}

	if plot {
		fmt.Println("Plotting not implemented in this Go version")
	}

	return sortedColors, nil
}

func main() {
	// Get command line args for image_path
	imagePath := os.Args[1]
	//imagePath := "/Users/mryan/GolandProjects/GMMBasedColorExtraction/detection_car01.png"
	dominantColors, err := findDominantColors(imagePath, false, 0.01, 1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Dominant colors:")
	for _, color := range dominantColors {
		rgbPoint := [3]float64{color.R, color.G, color.B}
		colorName := classifyColor(rgbPoint, comprehensiveColorRanges)
		fmt.Printf("R: %.0f, G: %.0f, B: %.0f - Classified as: %s\n", color.R, color.G, color.B, colorName)
	}
}
