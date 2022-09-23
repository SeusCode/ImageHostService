package controllers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"restapi/src/pkg/config"
	"time"

	"github.com/google/uuid"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(config.GetMaxUploadSize())       // Parse max 5 files
	files := r.MultipartForm.File[config.GetImagesForm()] // Get files

	os.Mkdir(config.GetImagesDir(), 0777)

	// No cache headers set
	var epoch = time.Unix(0, 0).Format(time.RFC1123)

	var noCacheHeaders = map[string]string{
		"Expires":         epoch,
		"Cache-Control":   "no-cache, private, max-age=0",
		"Pragma":          "no-cache",
		"X-Accel-Expires": "0",
	}

	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}

	// Iterate files
	//TODO: Limit to a config.GetMaxFilesUpload()
	for _, file := range files {
		fileData, err := file.Open()

		// Validations
		if err != nil {
			fmt.Fprintf(w, "No se pudo leer el archivo")
			return
		}

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, fileData)
		if err != nil {
			fmt.Fprintf(w, "No se pudo copiar la info en el buffer")
			return
		}

		//TODO: Add mime types in config file as json object (array)
		filetype := http.DetectContentType(buf.Bytes())
		if filetype != "image/png" && filetype != "image/jpg" && filetype != "image/gif" && filetype != "image/bmp" {
			fmt.Fprintf(w, "Tipo de imagen invalido!")
			return
		}

		if file.Size > (int64(config.GetMaxFileSize())) {
			fmt.Fprintf(w, "Archivo demasiado pesado!")
			return
		}

		if len(file.Filename) < 3 {
			fmt.Fprintf(w, "Nombre de archivo demasiado corto!")
			return
		}

		if len(file.Filename) > 30 {
			fmt.Fprintf(w, "Nombre de archivo demasiado largo!")
			return
		}

		// Generate a uuid
		id := uuid.New().String()

		dir := fmt.Sprintf("images/%c/%c/%c", id[0], id[1], id[2])

		os.MkdirAll(dir, 0644)

		ioutil.WriteFile(fmt.Sprintf("%s/%s_%s", dir, id, file.Filename), buf.Bytes(), 0777)
	}
}
