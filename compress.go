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
	T  uint8
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

		domains:
			//iterate over all domain blocks
			for dy := range domainBlocks {
				for dx := range domainBlocks[dy] {
					for i := 0; i < 8; i++ {
						tempImg := domainBlocks[dy][dx]
						copy(tempImg.Pix, domainBlocks[dy][dx].Pix)
						tempImg.TransformImage(uint8(i))
						tempVal := CalcSquarredEuclideanDistance(rangeBlocks[ry][rx], tempImg)

						if tempVal <= 0.001 {
							encodings[ry][rx] = EncodingParams{
								S:  1,
								G:  0,
								T:  uint8(i),
								Dx: dx,
								Dy: dy,
							}
							break domains

							//only check for optimal s and g if the value is already somehow low
						} else if tempVal < errorValue {
							tempVal = CalcSquarredEuclideanDistance(rangeBlocks[ry][rx], domainBlocks[dy][dx])
							errorValue = tempVal
							encodings[ry][rx] = EncodingParams{
								S:  uint8(1),
								G:  uint8(0),
								T:  uint8(i),
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
