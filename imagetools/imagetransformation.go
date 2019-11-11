package imagetools

import (
	"image"
	"image/color"
)

//FlipImage flips the image vertically by providing a pointer to an image and returning normal image.Gray struct
func FlipImage(img *image.Gray) *image.Gray {
	var imgCopy *image.Gray = image.NewGray(img.Rect)
	imgCopy.Pix = make([]uint8, len(imgCopy.Pix))

	//making a deep copy of img.Pix
	for i := range img.Pix {
		imgCopy.Pix[i] = img.Pix[i]
	}

	var grayValue = color.Gray{Y: 0}
	for i := 0; i < len(img.Pix); i++ {
		var x = i % img.Stride
		var y = i / img.Stride
		grayValue = img.GrayAt(img.Stride-x, y)
		imgCopy.SetGray(x, y, grayValue)
	}
	return imgCopy
}

// RotateImage creates a deep copy of the passed image pointer, then rotates this copy by x degrees and returns the deep copy pointer
func RotateImage(img *image.Gray) *image.Gray {
	var imgCopy *image.Gray = image.NewGray(img.Rect)
	imgCopy.Pix = make([]uint8, len(imgCopy.Pix))

	//making a deep copy of img.Pix
	for i := range img.Pix {
		imgCopy.Pix[i] = img.Pix[i]
	}

	var grayValue = color.Gray{Y: 0}
	for i := 0; i < len(img.Pix); i++ {
		var x = i % img.Stride
		var y = i / img.Stride
		grayValue = img.GrayAt(img.Stride-x, y)
		imgCopy.SetGray(x, y, grayValue)
	}

	return imgCopy
}
