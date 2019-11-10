package main

import (
	"fmt"
	"fractalcompression/imagetools"
	"log"
	"os"
)

//Sierpinski Triangle
var transformation1 = imagetools.Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0,
	F: 0,
}

var transformation2 = imagetools.Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0.5,
	F: 0,
}

var transformation3 = imagetools.Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0,
	F: 0.5,
}

// Cube creation
/*var transformation1 = imagetools.Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0,
	F: 0,
}

var transformation2 = imagetools.Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0,
	F: 0.5,
}

var transformation3 = imagetools.Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0.5,
	F: 0,
}

var transformation4 = imagetools.Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0.5,
	F: 0.5,
}*/

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

	grayImg = imagetools.CreateFractalFromImage(grayImg, 10, []imagetools.Transformation{transformation1, transformation2, transformation3})

	fmt.Println("Image successfully fractalized...")

	imagetools.SaveImageToFile(grayImg, filename)

	fmt.Println("Image successfully saved...")
}
