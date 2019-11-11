package imagetools

import (
	"image"
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
		img = applyIFSToImage(img, transformations)
	}
	return img
}

func applyIFSToImage(img *image.Gray, transformations []Transformation) *image.Gray {

	// (a b)  * (x)  + (e)
	// (c d)  * (y)  + (f)

	//create deep copy of img
	imgCopy := image.NewGray(img.Rect)
	for i := 0; i < len(imgCopy.Pix); i++ {
		imgCopy.Pix[i] = 255
	}

	for i := 0; i < len(transformations); i++ {
		for j := 0; j < len(imgCopy.Pix); j++ {
			var x = j % img.Stride
			var y = j / img.Stride
			grayValue := img.GrayAt(x, y)

			var newX int = int(math.Round(transformations[i].A*float64(x) + transformations[i].B*float64(y) + transformations[i].E*float64(img.Rect.Dx())))
			var newY int = int(math.Round(transformations[i].C*float64(x) + transformations[i].D*float64(y) + transformations[i].F*float64(img.Rect.Dy())))
			imgCopy.SetGray(newX, newY, grayValue)
		}
	}
	return imgCopy
}
