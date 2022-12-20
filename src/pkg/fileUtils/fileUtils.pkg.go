package fileUtils

import (
	"image"
	"io"
	"mime"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func FileSize(fPath string) (int64, error) {
	fi, err := os.Stat(fPath)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}

// Guess image format from gif/jpeg/png/webp
func GuessImageFormat(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	return
}

// Guess image mime types from gif/jpeg/png/webp
func GuessImageMimeTypes(r io.Reader) string {
	format, _ := GuessImageFormat(r)
	if format == "" {
		return ""
	}
	return mime.TypeByExtension("." + format)
}
