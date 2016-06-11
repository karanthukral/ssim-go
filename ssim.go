package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

// Default SSIM constants
var (
	L  = 255.0
	K1 = 0.01
	K2 = 0.03
	C1 = math.Pow((K1 * L), 2.0)
	C2 = math.Pow((K2 * L), 2.0)
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dimensions(img image.Image) (width, height int) {
	bounds := img.Bounds()
	width, height = bounds.Max.X, bounds.Max.y
	return
}

func getPixelValue(c color.Color) float64 {
	r, _, _, _ := c.RGBA()
	return float64(r >> 8)
}

func equalDim(img1, img2 image.Image) bool {
	w1, h1 := dimensions(img1)
	w2, h2 := dimensions(img2)
	return (w1 == w2) && (h1 == h2)
}

func convertToGrayscale(img image.Image) image.Image {
	imageBounds := img.Bounds()
	width, height := dimensions(img)

	grayImage := image.NewGray(imageBounds)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImage.Set(x, y, grayColor)
		}
	}

	return grayImage
}

func mean(img image.Image) float64 {
	width, height := dimensions(img)
	n := float64((width * height) - 1)
	sum := 0.0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			sum += getPixelValue(img.At(x, y))
		}
	}

	return sum / n
}

func stdDev(img image.Image) float64 {
	width, height := dimensions(img)

	n := float64((width * height) - 1)
	sum := 0.0
	avg := mean(img)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixel = getPixelValue(img.At(x, y))
			sum += math.Pow((pixel - avg), 2.0)
		}
	}

	return math.Sqrt(sum / n)
}

func covar(img1, img2 image.Image) (covar float64, err error) {
	if !equalDim(img1, img2) {
		err = errors.New("Images must have the same dimension")
		return
	}

	avg1 := mean(img1)
	avg2 := mean(img2)

	width, height := dimensions(img1)
	sum := 0
	n := float64((width * height) - 1)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pix1 := getPixelValue(img1.At(x, y))
			pix2 := getPixelValue(img2.At(x, y))
			sum += (pix1 - avg1) * (pix2 - avg2)
		}
	}

	c = sum / n
	return
}

func ssim(x, y image.Image) float64 {
	avg_x := mean(x)
	avg_y := mean(y)

	stdDev_x := stdDev(x)
	stdDev_y := stdDev(y)

	covar, err := covar(x, y)
	handleErr(err)

	num := ((2.0 * avg_x * avg_y) + C1) * ((2.0 * covar) + C2)
	denominator := (math.Pow(avg_x, 2.0) + math.Pow(avg_y, 2.0) + C1) *
		(math.Pow(stdev_x, 2.0) + math.Pow(stdev_y, 2.0) + C2)
	return num / denominator
}
