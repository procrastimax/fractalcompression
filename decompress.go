package main

import "image"

//DecodeImage .
func DecodeImage(encoding [][]EncodingParams, rangePixelSize, iterations int) *GrayImage {
	var img = NewGrayImage(image.Rect(0, 0, len(encoding)*rangePixelSize, len(encoding[0])*rangePixelSize))
	for it := 0; it < iterations; it++ {
		applyEncodings(encoding, img, rangePixelSize)
	}
	return img
}

//DecompressImage creates a gray image from a 2d array of encodings
// ranges are typically 4x4 pixel size
// requires an empty image
func applyEncodings(encoding [][]EncodingParams, img *GrayImage, pixelSize int) {
	ranges := img.DivideImage(pixelSize)
	// iterating over the ranges should be same as iterating over the encoding param 2d array
	for i := range ranges {
		for j := range ranges[i] {
			ranges[i][j].GrayTransformImageInPlace(Sarr[encoding[i][j].S], 0)
		}
	}
}
