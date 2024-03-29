package imagetools

import (
	"fmt"
	"image"
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
