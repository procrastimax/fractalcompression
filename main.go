package main

import (
	"fmt"
	"fractalcompression/imagetools"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please specify an imagepath!")
	}
	filename := os.Args[1]

	fmt.Println("Loading image and convert to fractal...")
	img := imagetools.LoadImageFromFile(filename)

	fmt.Println("Loading successfull...")

	grayImg := imagetools.ImageToGray(img)

	fmt.Println("Image successfully turned gray...")

	grayImg = imagetools.CreateFractalFromImage(grayImg)

	fmt.Println("Image successfully fractalized...")

	imagetools.SaveImageToFile(grayImg, filename)

	fmt.Println("Image successfully saved...")

}
