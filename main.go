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

	grayImg = imagetools.ScaleImage(grayImg, 0.2)
	//grayImg = imagetools.TransformImage(grayImg, 7)
	//grayImg = imagetools.CreateFractalFromImage(grayImg, 1, []imagetools.Transformation{transformation1, transformation2, transformation3})

	//fmt.Println("Creating range and domain arrays...")
	//ranges := imagetools.DivideImage(grayImg, 4)
	//domains := imagetools.DivideImage(grayImg, 8)

	/*fmt.Println("Scale down all domains to range-block size...")
	for i := range domains {
		for _, value := range domains[i] {
			value = imagetools.ScaleImage(value, 0.5)
		}
	}*/

	imagetools.SaveImageToFile(grayImg, filename)
	fmt.Println("Image successfully saved...")
}

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
