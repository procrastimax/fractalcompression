package imagetools

import (
	"image"
	"image/color"
	"math"
)

const iterations = 1

// Transformation is a struct to describe a ISF transformation
type Transformation struct {
	A, B, C, D, E, F float64
}

// CreateFractalFromImage is a Function to create fractal from ISF
func CreateFractalFromImage(img *image.Gray, numberOfIterations int, transformations []Transformation) *image.Gray {

	for i := 0; i < numberOfIterations; i++ {
		applyIFSToImage(img, transformations)
	}
	return img
}

func applyIFSToImage(img *image.Gray, transformations []Transformation) *image.Gray {

	// (a b)  * (x)  + (e)
	// (c d)  * (y)  + (f)

	imgMatrix := make([][]uint8, img.Stride)
	for i := 0; i < len(imgMatrix); i++ {
		imgMatrix[i] = make([]uint8, len(img.Pix)/img.Stride)
	}

	//fill img Matrix with 0 values
	var x, y int = 0, 0
	for i := 0; i < len(img.Pix); i++ {
		imgMatrix[x%img.Stride][y] = img.Pix[i]
		//make all pixel blank
		img.Pix[i] = 255
		x++
		if x%img.Stride == 0 {
			y++
		}
	}

	for i := 0; i < len(transformations); i++ {
		for x := 0; x < len(imgMatrix); x++ {
			for y := 0; y < len(imgMatrix[x]); y++ {
				grayValue := imgMatrix[x][y]
				var newX int = int(math.Round(transformations[i].A*float64(x) + transformations[i].B*float64(y) + transformations[i].E*float64(img.Rect.Dx())))
				var newY int = int(math.Round(transformations[i].C*float64(x) + transformations[i].D*float64(y) + transformations[i].F*float64(img.Rect.Dy())))
				img.SetGray(newX, newY, color.Gray{Y: grayValue})
			}
		}
	}

	return img
}
