package imagetools

import (
	"image"
	"image/color"
	"math"
)

const iterations = 1

func CreateFractalFromImage(img *image.Gray) *image.Gray {
	for i := 0; i < iterations; i++ {
		img = applyIFSToImage(img, 0.5, 0, 0, 0.5, 0, 0)
		img = applyIFSToImage(img, 0.5, 0, 0, 0.5, 0.5, 0)
		img = applyIFSToImage(img, 0.5, 0, 0, 0.5, 0, 0.5)
	}
	return img
}

func applyIFSToImage(img *image.Gray, a float64, b float64, c float64, d float64, e float64, f float64) *image.Gray {

	// a b  * x  + e
	// c d  * y  + f

	rect := img.Bounds()

	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			grayValue := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			//reset color at visited pixel
			//img.SetGray(x, y, color.Gray{255})
			//set this gray value on a transformed position
			var newX int = int(math.Round(a*float64(x) + b*float64(y) + e*float64(rect.Dx())))
			var newY int = int(math.Round(c*float64(x) + d*float64(y) + f*float64(rect.Dy())))

			img.SetGray(newX, newY, grayValue)
		}
	}
	return img
}
