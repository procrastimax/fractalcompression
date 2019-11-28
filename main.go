package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please specify an imagepath!")
	}
	filename := os.Args[1]

	fmt.Println("Loading image...")
	img := LoadImageFromFile(filename)
	fmt.Println("Loading successfull...")
	grayImg := ImageToGray(img)

	//domains := grayImg.DivideImage(8)
	//domains[1][1].SetGrayAtRelative(2, 2, 23)
	//fmt.Println(domains[1][1].GrayAtRelative(2, 2))

	encodingParams := EncodeImage(grayImg)
	*grayImg = *DecodeImage(encodingParams, 4, 10)

	fmt.Println("Saving image...")
	SaveImageToFileAsPNG(grayImg, filename)
}

//Sierpinski Triangle
var transformation1 = Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0,
	F: 0,
}

var transformation2 = Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0.5,
	F: 0,
}

var transformation3 = Transformation{
	A: 0.5,
	B: 0,
	C: 0,
	D: 0.5,
	E: 0,
	F: 0.5,
}
