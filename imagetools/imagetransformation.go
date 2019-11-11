package imagetools

import (
	"image"
	"image/color"
	"log"
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

// RotateImage creates a deep copy of the passed image pointer, then rotates this copy by (0)0,(1)90,(2)180,(3)270 degrees and then returns the rotated image
func RotateImage(img *image.Gray, rotationType uint8) *image.Gray {
	if rotationType < 0 || rotationType > 3 {
		log.Fatalln("Rotation type can only be a number between 0 and 3 representing 4 rotations (0,90,180,270)!")
	}
	var imgCopy *image.Gray = image.NewGray(img.Rect)
	imgCopy.Pix = make([]uint8, len(imgCopy.Pix))

	//making a deep copy of img.Pix
	for i := range img.Pix {
		imgCopy.Pix[i] = 0
	}

	var grayValue = color.Gray{Y: 0}
	for i := 0; i < len(img.Pix); i++ {
		var x = i % img.Stride
		var y = i / img.Stride

		// add sin/cos calculation here
		var newX = 0
		var newY = 0

		switch rotationType {
		case 0:
			{
				newX = x
				newY = y
			}
		case 1:
			{
				newX = img.Rect.Dx() - y
				newY = x
			}
		case 2:
			{
				newX = img.Rect.Dx() - x
				newY = img.Rect.Dy() - y
			}
		case 3:
			{
				newX = y
				newY = img.Rect.Dy() - x
			}
		}

		grayValue = img.GrayAt(x, y)
		imgCopy.SetGray(newX, newY, grayValue)
	}

	return imgCopy
}

// DivideImage slices an image into pixelSize*pixelSize smaller images and returning them in an array
func DivideImage(img *image.Gray, pixelSize int) [][]*image.Gray {

	// how many new smaller images are being created per axis
	var dividedCount int = img.Rect.Dx() / int(pixelSize)

	imageParts := make([][]*image.Gray, dividedCount)

	for i := 0; i < dividedCount; i++ {
		imageParts[i] = make([]*image.Gray, dividedCount)
		for j := 0; j < dividedCount; j++ {
			rect := image.Rect(i*pixelSize, j*pixelSize, (pixelSize*i)+pixelSize, (pixelSize*j)+pixelSize)
			imageParts[i][j] = img.SubImage(rect).(*image.Gray)
		}
	}
	return imageParts
}
