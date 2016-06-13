package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"reflect"
)

func main() {
	file, err := os.Open("pic.jpg")
	if err != nil {
		log.Fatalf("Could Not Open Pic -> %v", err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Could Not Decode Pic -> %v", err)
	}

	file, err = os.Open("pic2.jpg")
	if err != nil {
		log.Fatalf("Could Not Open Pic2 -> %v", err)
	}

	img2, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Could Not Decode Pic2 -> %v", err)
	}
	index := ssim(img, img2)
	fmt.Printf("Index is type of = %v\n", reflect.TypeOf(index))
	fmt.Printf("Index Value = %f", index)
}
