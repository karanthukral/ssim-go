package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	file, err := os.Open("pic.jpg")
	if err != nil {
		log.Fatalf("Shit broke")
	}

	img, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Shit not good")
	}
	index := ssim(x, y)
	print(index)
}
