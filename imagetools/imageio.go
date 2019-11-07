package imagetools

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
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

func SaveImageToFile(img *image.Gray, filename string) {
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
