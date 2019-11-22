package imagetools

import (
	"image"
)

//Decompress sad
func Decompress(encoding [][]EncodingParams, iterations int) *image.Gray {
	var img = image.NewGray(image.Rect(0, 0, len(encoding)*4, len(encoding[0])*4))
	for i := range img.Pix {
		img.Pix[i] = 255
	}
	for it := 0; it < iterations; it++ {
		applyEncodings(encoding, img)
	}
	return img
}

//DecompressImage creates a gray image from a 2d array of encodings
// ranges are typically 4x4 pixel size
// requires an empty image
func applyEncodings(encoding [][]EncodingParams, img *image.Gray) {
	ranges := DivideImage(img, 4)
	// iterating over the ranges should be same as iterating over the encoding param 2d array
	for i := range ranges {
		for j := range ranges[i] {
			*ranges[i][j] = *GrayTransformImage(ranges[i][j], encoding[i][j].S, encoding[i][j].G)
		}
	}

	// create a whole image from the ranges
	for xR := range ranges {
		for yR := range ranges[xR] {
			for i := range ranges[xR][yR].Pix {
				tempX := i % ranges[xR][yR].Stride
				tempY := i / ranges[xR][yR].Stride
				img.Pix[img.PixOffset((xR*4)+tempX, (yR*4)+tempY)] = ranges[xR][yR].Pix[i]
			}
		}
	}
}
