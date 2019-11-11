package imagetools

import (
	"image"
	"image/color"
)

//ImageToGray takes an image.Image pointer and returns it as image.Gray
func ImageToGray(img *image.Image) *image.Gray {
	size := (*img).Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	newImg := image.NewGray(rect)
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := (*img).At(x, y)
			newImg.SetGray(x, y, color.GrayModel.Convert(pixel).(color.Gray))
		}
	}
	return newImg
}

//ImageToBW takes an image.Image pointer and return it as image.Gray where only pixelvalues of 0 and 255 are allowed (only black and white)
func ImageToBW(img *image.Image) *image.Gray {
	size := (*img).Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	newImg := image.NewGray(rect)
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := (*img).At(x, y)
			bwValue := color.GrayModel.Convert(pixel).(color.Gray)

			if bwValue.Y >= 128 {
				bwValue.Y = 255
			} else {
				bwValue.Y = 0
			}
			newImg.SetGray(x, y, bwValue)
		}
	}
	return newImg
}
