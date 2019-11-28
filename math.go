package main

import (
	"fmt"
	"log"
	"math"
)

//Sarr the quantized parameters for s
var Sarr = [...]float64{0, 0.05, 0.1, 0.25, 0.5, 1, 1.5, 2}

//Garr the quantized parameters for g
var Garr = [...]float64{-5, -2, -1, 0, 1, 2, 5, 10}

//CalcSquarredEuclideanDistance calculates the euclidean distance between a range and a domain block
func CalcSquarredEuclideanDistance(rangeBlock *GrayImage, domainBlock *GrayImage) float64 {
	if rangeBlock.Bounds().Dx() != domainBlock.Bounds().Dx() {
		fmt.Println(rangeBlock.Rect, domainBlock.Rect)
		log.Fatalln("CalcSquarredEuclideanDistance: Rects of domain and range block do not share same width!")
	}
	var errorValue = 0.0

	rBounds := rangeBlock.Bounds()
	for y := 0; y < rBounds.Dy(); y++ {
		for x := 0; x < rBounds.Dy(); x++ {
			//fmt.Println(rangeBlock.Bounds(), domainBlock.Bounds(), x, y, x-rBounds.Min.X+dBounds.Min.X, y-rBounds.Min.Y+dBounds.Min.Y)
			errorValue += math.Pow(float64(domainBlock.GrayAtRelative(x, y)-rangeBlock.GrayAtRelative(x, y)), 2.0)
		}
	}
	return math.Sqrt(errorValue)
}