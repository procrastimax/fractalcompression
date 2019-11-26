package main

import (
	"fmt"
	"math"
)

//EncodingParams saves the encoding parameters for every range which shall be recreated
type EncodingParams struct {
	S  uint8
	G  uint8
	Dx int
	Dy int
}

//EncodeImage encodes the image and returns encoding params as a 2D array
func EncodeImage(img *GrayImage) [][]EncodingParams {
	rangeBlocks := img.DivideImage(4)
	domainBlocks := img.DivideImage(8)

	for i := range domainBlocks {
		for j := range domainBlocks[i] {
			domainBlocks[i][j].ScaleImage(0.5)
		}
	}

	return findBestMatchingDomains(rangeBlocks, domainBlocks)
}

//FindBestMatchingDomains finds the best matching domain
func findBestMatchingDomains(rangeBlocks [][]*GrayImage, domainBlocks [][]*GrayImage) [][]EncodingParams {
	var encodings = make([][]EncodingParams, len(rangeBlocks))

	//iterate over all range blocks
	for ry := range rangeBlocks {
		fmt.Printf("%d from %d\n", ry, len(rangeBlocks)-1)
		encodings[ry] = make([]EncodingParams, len(rangeBlocks[ry]))
		for rx := range rangeBlocks[ry] {
			errorValue := math.MaxFloat64
			//iterate over all domain blocks
			for dy := range domainBlocks {
				for dx := range domainBlocks[dy] {
					for sIT, s := range Sarr {
						tempImg := domainBlocks[dy][dx].GrayTransformImage(s, 0)
						tempVal := CalcSquarredEuclideanDistance(rangeBlocks[ry][rx], tempImg)
						if tempVal < errorValue {
							errorValue = tempVal
							encodings[ry][rx] = EncodingParams{
								S:  uint8(sIT),
								G:  0,
								Dx: dx,
								Dy: dy,
							}
						}
					}
				}
			}
		}
	}
	return encodings
}
