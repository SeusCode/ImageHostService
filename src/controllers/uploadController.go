/*

 * Original Idea: Sergio Baudracco. First code: Genaro Civilotti
 * @author SEUSCODE
 * @date 19/SEPT/2022
 * @version 1.0
 * Copyright SEUSCODE 2022
 *
 * Reservados todos los derechos. Prohibida la copia y distribución de este programa y los recursos que lo acompañan.
 *
 */

package controllers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"restapi/src/libs"
	"restapi/src/models"
	"restapi/src/pkg/config"
	"time"

	"github.com/google/uuid"
)

func Upload(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Add("Server", "SeusCode-Images/1.0")
	w.Header().Add("Seus", "Images/1.0")
	//w.Header().Add("Cache-Control", "public, max-age=86400")

	r.ParseMultipartForm(config.GetConfig().MAX_UPLOAD_SIZE)      // Parse max 5 files
	files := r.MultipartForm.File[config.GetConfig().IMAGES_FORM] // Get files

	os.Mkdir(config.GetConfig().IMAGES_DIR, 0777)

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

	if len(files) > config.GetConfig().MAX_FILES_UPLOAD {
		fmt.Fprintf(w, "Demasiados archivos")

		w.WriteHeader(500)
		w.Write([]byte("too many files"))

		return
	}

	// Iterate files
	for _, file := range files {
		fileData, err := file.Open()

		// Validations
		if err != nil {
			fmt.Fprintf(w, "No se pudo leer el archivo")

			w.WriteHeader(500)
			w.Write([]byte("can't read file"))

			return
		}

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, fileData)
		if err != nil {
			fmt.Fprintf(w, "No se pudo copiar la info en el buffer")

			w.WriteHeader(500)
			w.Write([]byte("can't read file buffer"))

			return
		}

		//TODO: Add mime types in config file as json object (array)
		filetype := http.DetectContentType(buf.Bytes())
		if filetype != "image/webp" && filetype != "image/png" && filetype != "image/jpg" && filetype != "image/jpeg" && filetype != "image/gif" && filetype != "image/bmp" {
			fmt.Fprintf(w, "Tipo de imagen invalido!")
			fmt.Print(filetype)

			w.WriteHeader(500)
			w.Write([]byte("wrong filetype"))
			return
		}

		if file.Size > (int64(config.GetConfig().MAX_FILE_SIZE)) {
			fmt.Fprintf(w, "Archivo demasiado pesado!")

			w.WriteHeader(500)
			w.Write([]byte("wrong file size"))
			return
		}

		if len(file.Filename) < 3 {
			fmt.Fprintf(w, "Nombre de archivo demasiado corto!")

			w.WriteHeader(500)
			w.Write([]byte("wrong file name, too short"))

			return
		}

		if len(file.Filename) > 30 {
			fmt.Fprintf(w, "Nombre de archivo demasiado largo!")

			w.WriteHeader(500)
			w.Write([]byte("wrong file name, too large"))

			return
		}

		// Generate a uuid
		id := uuid.New().String()

		dir := fmt.Sprintf("images/%c/%c/%c", id[0], id[1], id[2])

		os.MkdirAll(dir, 0644)

		ioutil.WriteFile(fmt.Sprintf("%s/%s_%s", dir, id, file.Filename), buf.Bytes(), 0777)

		//msg := fmt.Sprintf("Imagen subida con exito.\n\nhttp://localhost:%d/%s/%s_%s", config.GetConfig().SERVER_PORT, dir, id, file.Filename)
		msg := fmt.Sprintf("/%s/%s_%s", dir, id, file.Filename)

		ip := r.RemoteAddr

		libs.DB.Create(&models.SystemLog{
			Name: fmt.Sprintf("Upload %s", id),
			Ip:   ip,
		})

		libs.DB.Create(&models.Image{
			Name: file.Filename,
			Uid:  id,
		})

		fmt.Fprintf(w, msg)

		w.WriteHeader(200)
		w.Write([]byte("upload ok"))
	}
}
