package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
)

//--------------------------
//Main Interface implementation
//--------------------------

//GrayImage represents a gray image with a 2D gray value matrix
type GrayImage struct {
	// Pix holds the image's pixels, as gray values. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
	Pix    []uint8
	Stride int
	Rect   image.Rectangle
}

//At returns the color value at pixel x,y
func (img *GrayImage) At(x, y int) color.Color {
	//theoretically, the max point isnt included inside the rectangle (just like in slices)
	//but this is a bit unlogically for me, so Im going to include the max point
	if img.Rect.Min.X <= x ||
		img.Rect.Max.X > x ||
		img.Rect.Min.Y <= y ||
		img.Rect.Max.Y > y {
		return color.Gray{img.Pix[(y-img.Rect.Min.Y)*img.Stride+x-(img.Rect.Min.X)]}
	}
	fmt.Println("At: Trying to get pixel value which is outside of the boundaries!")
	return color.Gray{0}
}

//GrayAt returns the grayvalue as an uint8 values at x,y
func (img *GrayImage) GrayAt(x, y int) uint8 {
	if img.Rect.Min.X <= x ||
		img.Rect.Max.X > x ||
		img.Rect.Min.Y <= y ||
		img.Rect.Max.Y > y {
		return img.Pix[(y-img.Rect.Min.Y)*img.Stride+x-(img.Rect.Min.X)]
	}
	fmt.Println("GrayAt: Trying to get pixel value which is outside of the boundaries!")
	return 0
}

//GrayAtRelative returns the grayvalue as an uint8 values at x,y which are relative values,
//so the first pixel would x = 0, y = 0
func (img *GrayImage) GrayAtRelative(x, y int) uint8 {

	xRel := x + img.Bounds().Min.X
	yRel := y + img.Bounds().Min.Y
	if img.Rect.Min.X <= xRel ||
		img.Rect.Max.X > xRel ||
		img.Rect.Min.Y <= yRel ||
		img.Rect.Max.Y > yRel {
		return img.Pix[(yRel-img.Rect.Min.Y)*img.Stride+xRel-(img.Rect.Min.X)]
	}
	fmt.Println("GrayAt: Trying to get pixel value which is outside of the boundaries!")
	return 0
}

//ColorModel returns the image's color model
func (img *GrayImage) ColorModel() color.Model {
	return color.GrayModel
}

//Bounds returns the domain for which At can return non-zero color values
func (img *GrayImage) Bounds() image.Rectangle {
	return img.Rect
}

//NewGrayImage creates an empty new grayImage from a given rectangle
func NewGrayImage(rect image.Rectangle) *GrayImage {
	pix := make([]uint8, rect.Dx()*rect.Dy())
	for i := range pix {
		pix[i] = 0
	}
	return &GrayImage{
		Pix:    pix,
		Stride: rect.Dx(),
		Rect:   rect,
	}
}

//SetGrayAt sets a uint8 gray value at position x,y
//This function respects the boundaries
func (img *GrayImage) SetGrayAt(x, y int, grayValue uint8) {
	if !(image.Point{x, y}.In(img.Rect)) {
		fmt.Println(img.Rect.String() + " - P: " + image.Point{x, y}.String() + " ")
		log.Fatalln("SetGrayAt: Trying to set pixel value which is outside of the boundaries!")
		return
	}
	i := img.PixOffset(x, y)
	img.Pix[i] = grayValue
}

//SetGrayAtRelative sets the gray value at relative pixel x,y, the first pixel would be x = 0, y = 0
func (img *GrayImage) SetGrayAtRelative(x, y int, grayvalue uint8) {

	xRel := x + img.Bounds().Min.X
	yRel := y + img.Bounds().Min.Y
	if !(image.Point{xRel, yRel}.In(img.Rect)) {
		fmt.Println(img.Rect.String() + " - P: " + image.Point{xRel, yRel}.String() + " ")
		log.Fatalln("SetGrayAtRelative: Trying to set pixel value which is outside of the boundaries!")
		return
	}
	i := img.PixOffset(xRel, yRel)
	img.Pix[i] = grayvalue
}

//PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x,y)
func (img *GrayImage) PixOffset(x, y int) int {
	return (y-img.Rect.Min.Y)*img.Stride + (x - img.Rect.Min.X)
}

//SubImage returns the same basic image with changed boundaries
func (img *GrayImage) SubImage(rect image.Rectangle) *GrayImage {
	rect = rect.Intersect(img.Rect)
	if rect.Empty() {
		fmt.Println("SubImage: Rectangles doesn't intersect each other, returning empty GrayImage")
		return &GrayImage{}
	}
	i := img.PixOffset(rect.Min.X, rect.Min.Y)
	return &GrayImage{
		Pix:    img.Pix[i:],
		Stride: img.Stride,
		Rect:   rect,
	}
}

//--------------------------
//Extended Functions
//--------------------------

//CalcAverageGrayLevel iterates over all pixel in the image and calcuates the average gray level
//works like At, so it only calculates the gray value from the values inside the boundaries
func (img *GrayImage) CalcAverageGrayLevel() uint8 {
	b := img.Bounds()
	grayValue, counter := 0, 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			grayValue += int(img.GrayAt(x, y))
			counter++
		}
	}
	return uint8(grayValue / counter)
}

// Transformation is a struct to describe a ISF transformation
type Transformation struct {
	A, B, C, D, E, F float64
}

//applyTransformationToGrayImage applies an affine transformation to the grayimage
//The transformation is only applied to the valid boundaries of the image
func (img *GrayImage) applyTransformationToGrayImage(transformation Transformation) {
	// (a b)  * (x)  + (e)
	// (c d)  * (y)  + (f)
	pixCopy := make([]uint8, len(img.Pix))
	copy(pixCopy, img.Pix)

	b := img.Bounds()
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			grayValue := pixCopy[img.PixOffset(x, y)]
			newX := int(math.Round(transformation.A*float64(x) + transformation.B*float64(y) + transformation.E*float64(b.Size().X)))
			newY := int(math.Round(transformation.C*float64(x) + transformation.D*float64(y) + transformation.F*float64(b.Size().Y)))
			img.SetGrayAtRelative(newX, newY, grayValue)
		}
	}
}

// TransformImage applies one of 8 transformations to the image
// 0 - Identity
// 1 - Rotation by 90°
// 2 - Rotation by 180°
// 3 - Rotation by 270°
// 4 - Flip at y axis
// 5 - Flip at x axis
// 6 - Mirror at y=x
// 7 - Mirror at y=-x
func (img *GrayImage) TransformImage(transformationType uint8) {
	if transformationType < 0 || transformationType > 7 {
		log.Fatalln("Transform type can only be a number between 0 and 7 representing 4 rotations and their mirrors!")
	}
	var transformation Transformation

	switch transformationType {
	case 0:
		//is identity, we can return the given image
		{
			return
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
	img.applyTransformationToGrayImage(transformation)
}

// DivideImage slices an image into pixelSize*pixelSize smaller images and returning them in an array
func (img *GrayImage) DivideImage(pixelSize int) [][]*GrayImage {
	// how many new smaller images are being created per axis
	var dividedCount int = img.Rect.Dx() / int(pixelSize)

	imageParts := make([][]*GrayImage, dividedCount)

	for i := 0; i < dividedCount; i++ {
		imageParts[i] = make([]*GrayImage, dividedCount)
		for j := 0; j < dividedCount; j++ {
			imageParts[i][j] = img.SubImage(image.Rect(
				i*pixelSize+img.Bounds().Min.X,
				j*pixelSize+img.Bounds().Min.Y,
				i*pixelSize+img.Bounds().Min.X+pixelSize,
				j*pixelSize+img.Bounds().Min.Y+pixelSize,
			))
		}
	}
	return imageParts
}

//ScaleImage scales a given image pointer by *scalingFactor* and returns the result
//Use only scalingFactor <= 1 and > 0!
func (img *GrayImage) ScaleImage(scalingFactor float64) {
	if scalingFactor > 0.5 || scalingFactor <= 0 {
		log.Fatalln("Scaling Factor has to be greater than 0 and less or equal than 0.5! Scaling values abore 0.5 are currently not implemented...")
	}
	oldWidth := img.Bounds().Dx()
	//create smaller subimage from an 8x8 image to an 4x4 when scaling factor is 0.5
	newRect := image.Rect(
		img.Rect.Min.X,
		img.Rect.Min.Y,
		img.Rect.Min.X+int(float64(img.Bounds().Dx())*scalingFactor),
		img.Rect.Min.Y+int(float64(img.Bounds().Dy())*scalingFactor))

	newWidth := newRect.Bounds().Dx()

	//newPixelBlockWidth is the value of pixel width of for the new blocks which are going to be averaged
	newPixelBlockWidth := oldWidth / newWidth

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y += newPixelBlockWidth {
		for x := b.Min.X; x < b.Max.X; x += newPixelBlockWidth {
			grayValue := 0
			for a := y; a < y+newPixelBlockWidth; a++ {
				for b := x; b < x+newPixelBlockWidth; b++ {
					grayValue += int(img.GrayAt(x, y))
				}
			}
			grayValue = grayValue / (newPixelBlockWidth * newPixelBlockWidth)
			img.SetGrayAt(((x + b.Min.X) / newPixelBlockWidth), ((y + b.Min.Y) / newPixelBlockWidth), uint8(grayValue))
		}
	}
	*img = *img.SubImage(newRect)
}

//GrayTransformImage applies contrast and brightness transformation to image
func (img *GrayImage) GrayTransformImage(s float64, g float64) *GrayImage {

	pixCopy := make([]uint8, len(img.Pix))
	copy(pixCopy, img.Pix)

	imgCopy := &GrayImage{pixCopy, img.Stride, img.Rect}

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			newGray := int((s * float64(img.GrayAt(x, y))) + g)
			if newGray > 255 {
				imgCopy.SetGrayAt(x, y, 255)
			} else if newGray < 0 {
				imgCopy.SetGrayAt(x, y, 0)
			} else {
				imgCopy.SetGrayAt(x, y, uint8(newGray))
			}
		}
	}
	return imgCopy
}

//GrayTransformImageInPlace applies contrast and brightness transformation to image
func (img *GrayImage) GrayTransformImageInPlace(s float64, g float64) {

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			newGray := int((s * float64(img.GrayAt(x, y))) + g)
			if newGray > 255 {
				img.SetGrayAt(x, y, 255)
			} else if newGray < 0 {
				img.SetGrayAt(x, y, 0)
			} else {
				img.SetGrayAt(x, y, uint8(newGray))
			}
		}
	}
}

//ScaleImage2 scales via an ifs
//is ok for minimal scaling, but doesnt do well on more extreme scalings
func (img *GrayImage) ScaleImage2(scalingFactor float64) {
	var transformation = Transformation{
		A: 1.0 * scalingFactor,
		B: 0.0,
		C: 0.0,
		D: 1.0 * scalingFactor,
		E: 0.0,
		F: 0.0,
	}
	newRect := image.Rect(
		img.Rect.Min.X,
		img.Rect.Min.Y,
		img.Rect.Min.X+int(float64(img.Bounds().Dx())*scalingFactor),
		img.Rect.Min.Y+int(float64(img.Bounds().Dy())*scalingFactor))

	img.applyTransformationToGrayImage(transformation)

	*img = *img.SubImage(newRect)
}
