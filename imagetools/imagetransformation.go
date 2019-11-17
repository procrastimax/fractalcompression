package imagetools

import (
	"image"
	"image/color"
	"log"
)

// TransformImage applies one of 8 transformations to the image
// 0 - Identity
// 1 - Rotation by 90°
// 2 - Rotation by 180°
// 3 - Rotation by 270°
// 4 - Flip at y axis
// 5 - Flip at x axis
// 6 - Mirror at y=x
// 7 - Mirror at y=-x
func TransformImage(img *image.Gray, rotationType uint8) *image.Gray {
	if rotationType < 0 || rotationType > 7 {
		log.Fatalln("Transform type can only be a number between 0 and 3 representing 4 rotations (0,90,180,270)!")
	}
	var transformation Transformation

	switch rotationType {
	case 0:
		//is identity, we can return the given image
		{
			return img
		}
	case 1:
		//rotation by 90°
		{
			transformation = Transformation{
				A: 0.0,
				B: 1.0,
				C: -1.0,
				D: 0.0,
				E: 0.0,
				F: 1.0,
			}
		}
	case 2:
		//rotation by 180°
		{
			{
				transformation = Transformation{
					A: -1.0,
					B: 0.0,
					C: 0.0,
					D: -1.0,
					E: 1.0,
					F: 1.0,
				}
			}
		}
	case 3:
		//rotation by 270°
		{
			{
				transformation = Transformation{
					A: 0.0,
					B: -1.0,
					C: 1.0,
					D: 0.0,
					E: 1.0,
					F: 0.0,
				}
			}
		}
	case 4:
		// Flip at y axis
		{
			transformation = Transformation{
				A: -1.0,
				B: 0.0,
				C: 0.0,
				D: 1.0,
				E: 1.0,
				F: 0.0,
			}
		}
	case 5:
		// Flip at x axis
		{
			transformation = Transformation{
				A: 1.0,
				B: 0.0,
				C: 0.0,
				D: -1.0,
				E: 0.0,
				F: 1.0,
			}
		}
	case 6:
		// Mirror at y=x
		{
			transformation = Transformation{
				A: 0.0,
				B: 1.0,
				C: 1.0,
				D: 0.0,
				E: 0.0,
				F: 0.0,
			}
		}
	case 7:
		// Mirror at y=-x
		{
			transformation = Transformation{
				A: 0.0,
				B: -1.0,
				C: -1.0,
				D: 0.0,
				E: 1.0,
				F: 1.0,
			}
		}
	}
	return applyIFSToImage(img, []Transformation{transformation})
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
