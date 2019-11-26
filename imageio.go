package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// LoadImageFromFile returns a pointer to an image.Image by providing a filename
func LoadImageFromFile(filename string) *image.Image {
	if len(strings.TrimSpace(filename)) == 0 {
		log.Fatalln("Image path shall not be null or empty!")
	}

	f, err := os.Open(filename)
	check(err)
	defer f.Close()
	img, format, err := image.Decode(f)
	check(err)
	if format != "jpeg" {
		log.Fatalln("Only jpeg images are supported!")
	}
	return &img
}

// SaveImageToFile saves a i*mg.Gray pointer to the specified filename with the *_edited* filepostix
func SaveImageToFile(img *GrayImage, filename string) {
	if len(strings.TrimSpace(filename)) == 0 {
		log.Fatalln("Image path shall not be null or empty!")
	}
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filepath.Base(filename), ext)
	newImagePath := fmt.Sprintf("%s/%s_edited%s", filepath.Dir(filename), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, img, nil)
	check(err)
}

//ImageToGray takes an image.Image pointer and returns it as image.Gray
func ImageToGray(img *image.Image) *GrayImage {
	b := (*img).Bounds()
	newImg := NewGrayImage((*img).Bounds())
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pixel := (*img).At(x, y)
			newImg.SetGrayAt(x, y, color.GrayModel.Convert(pixel).(color.Gray).Y)
		}
	}
	return newImg
}

//ImageToBW takes an image.Image pointer and return it as image.Gray where only pixelvalues of 0 and 255 are allowed (only black and white)
func ImageToBW(img *image.Image) *GrayImage {
	b := (*img).Bounds()
	newImg := NewGrayImage((*img).Bounds())
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			grayValue := color.GrayModel.Convert((*img).At(x, y)).(color.Gray).Y

			if grayValue >= 125 {
				grayValue = 255
			} else {
				grayValue = 0
			}
			newImg.SetGrayAt(x, y, grayValue)
		}
	}
	return newImg
}
