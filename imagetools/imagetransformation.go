package imagetools

import (
	"image"
	"image/color"
	"log"
	"math"
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
func TransformImage(img *image.Gray, transformationType uint8) *image.Gray {
	if transformationType < 0 || transformationType > 7 {
		log.Fatalln("Transform type can only be a number between 0 and 3 representing 4 rotations (0,90,180,270)!")
	}
	var transformation Transformation

	switch transformationType {
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
			rect := image.Rect(0, 0, pixelSize, pixelSize)
			imageParts[i][j] = image.NewGray(rect)
			for c := range imageParts[i][j].Pix {
				grayValue := img.GrayAt(i*pixelSize+c, j*pixelSize+c)
				imageParts[i][j].Pix[c] = grayValue.Y
			}
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

	for i := range imgCopy.Pix {
		imgCopy.Pix[i] = 0
	}

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
	return imgCopy
}

//GrayTransformImage applies contrast and brightness transformation to image
func GrayTransformImage(img *image.Gray, s float64, g float64) *image.Gray {
	for i := range img.Pix {
		grayValue := int((s*1.5)*float64(img.Pix[i]) + (g * 20))
		if grayValue < 0 {
			img.Pix[i] = 0
		} else if grayValue > 255 {
			img.Pix[i] = 255
		} else {
			img.Pix[i] = uint8(grayValue)
		}
	}
	return img
}

//ScaleImage2 scales via an ifs
//is ok for minimal scaling, but doesnt do well on more extreme scalings
func ScaleImage2(img *image.Gray, scalingFactor float64) *image.Gray {
	var transformation = Transformation{
		A: 1.0 * scalingFactor,
		B: 0.0,
		C: 0.0,
		D: 1.0 * scalingFactor,
		E: 0.0,
		F: 0.0,
	}
	img = applyIFSToImage(img, []Transformation{transformation})
	img.Rect = image.Rect(0, 0, int(float64(img.Bounds().Dx())*scalingFactor), int(float64(img.Bounds().Dy())*scalingFactor))
	return img
}

//CalcSquarredEuclideanDistance calculates the euclidean distance between a range and a domain block
// it returns the euclidean distance, and the parameters s and g
func CalcSquarredEuclideanDistance(rangeBlock *image.Gray, domainBlock *image.Gray) (float64, float64, float64) {
	s, g := calcContrastAndBrightness(rangeBlock, domainBlock)
	var errorValue = 0.0
	for i := range rangeBlock.Pix {
		errorValue += math.Pow((s*float64(domainBlock.Pix[i])+g)-float64(rangeBlock.Pix[i]), 2.0)
	}
	return errorValue, s, g
}

//calcContrast using this function: http://einstein.informatik.uni-oldenburg.de/rechnernetze/fraktal.htm
func calcContrastAndBrightness(rangeBlock *image.Gray, domainBlock *image.Gray) (float64, float64) {
	//calculating s - contrast
	var n = float64(len(rangeBlock.Pix))
	var numeratorS = 0.0

	var multRD = 0.0
	var sumR = 0.0
	var sumD = 0.0
	var sumD2 = 0.0
	for i := range rangeBlock.Pix {
		multRD += float64(rangeBlock.Pix[i] * domainBlock.Pix[i])
		sumR += float64(rangeBlock.Pix[i])
		sumD += float64(domainBlock.Pix[i])
		sumD2 += math.Pow(float64(domainBlock.Pix[i]), 2.0)
	}

	//fmt.Println(n, multRD, sumR, sumD, sumD2)

	numeratorS = ((n * n) * float64(multRD)) - float64(sumR*sumD)
	//fmt.Println("numeratorS " + strconv.FormatFloat(numeratorS, 'f', 5, 64))

	var denominatorS float64 = (n * n * float64(sumD2)) - math.Pow(float64(sumD), 2.0)

	//fmt.Println("denominatorS " + strconv.FormatFloat(denominatorS, 'f', 5, 64))
	var contrast = (numeratorS / denominatorS)

	//calculating g - brightness
	var numeratorO = float64(sumR) - (contrast * float64(sumD))
	var denominatorO = math.Pow(float64(n), 2.0)
	var brightness = numeratorO / denominatorO
	//fmt.Println("contrast, brightness "+strconv.FormatFloat(contrast, 'f', 5, 64), strconv.FormatFloat(brightness, 'f', 5, 64))
	return contrast, brightness
}
