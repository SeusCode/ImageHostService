package imageUtils

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Save save img
func ImageSave(name string, img image.Image) (err error) {
	var filetype = filepath.Ext(name)
	file, err := os.Create(name)
	defer file.Close()
	switch filetype {
	case ".jpeg", ".jpg":
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			return err
		}
	case ".gif":
		err = gif.Encode(file, img, nil)
		if err != nil {
			return err
		}
	case ".png":
		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	case ".webp":
		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	case ".bmp":
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			return err
		}
	default:
		return errors.New("type no match")
	}
	return nil
}
