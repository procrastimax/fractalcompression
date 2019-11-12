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

//ScaleImage scales a given image pointer by *scalingFactor* and returns the result
//Use only scalingFactor <= 1 and > 0!
func ScaleImage(img *image.Gray, scalingFactor float64) *image.Gray {
	if scalingFactor > 0.5 || scalingFactor <= 0 {
		log.Fatalln("Scaling Factor has to be greater than 0 and less or equal than 0.5! Scaling values abore 0.5 are currently not implemented...")
	}

	//pixel averaging -> explanation see following link:
	//https://entropymine.com/imageworsener/pixelmixing/

	//tileCount is the new value of D.x() and D.y()
	tileCount := int(float64(img.Stride) * scalingFactor)

	// make deep copy of passed image
	imgCopy := image.NewGray(image.Rect(0, 0, tileCount, tileCount))
	imgCopy = img

	//pixelBox is the value of pixel inside a tile on the original image
	pixelCount := img.Stride / imgCopy.Stride

	var x = 0
	var y = 0
	var grayValue = color.Gray{Y: 255}
	var grayInt = 0

	//iterave over every pixel of the new image and set the gray value as an average value from the old image
	for i := range imgCopy.Pix {
		grayInt = 0
		x = (i % imgCopy.Stride)
		y = (i / imgCopy.Stride)

		for a := 0; a < pixelCount; a++ {
			for b := 0; b < pixelCount; b++ {
				grayInt += int(img.GrayAt((pixelCount*x)+a, (pixelCount*y)+b).Y)
			}
		}

		//average grayValue
		grayValue.Y = uint8(grayInt / (pixelCount * pixelCount))
		imgCopy.SetGray(x, y, grayValue)
	}
	return img
}
