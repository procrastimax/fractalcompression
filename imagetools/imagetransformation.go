package imagetools

import (
	"image"
	"image/color"
)

//FlipImage flips the image vertically by providing a pointer to an image and returning normal image.Gray struct
func FlipImage(img *image.Gray) image.Gray {
	var imgCopy image.Gray = *image.NewGray(img.Rect)
	imgCopy.Pix = make([]uint8, len(imgCopy.Pix))

	//making a deep copy of img.Pix
	for i := range img.Pix {
		imgCopy.Pix[i] = img.Pix[i]
	}

	var grayValue = color.Gray{Y: 0}
	for i := 0; i < len(img.Pix); i++ {
		grayValue = img.GrayAt(img.Stride-(i%img.Stride), i/img.Stride)
		imgCopy.SetGray(i%img.Stride, i/img.Stride, grayValue)
	}
	return imgCopy
}
