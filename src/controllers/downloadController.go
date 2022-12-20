/*
 * SEUS Code ©2022
 * @Author: SERBice
 */
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/sergeymakinen/go-bmp"
	"golang.org/x/image/webp"

	//Libraries tests
	//"github.com/kolesa-team/go-webp/webp"
	//"github.com/kolesa-team/go-webp/encoder"
	//"github.com/kolesa-team/go-webp/decoder"
	//"github.com/chai2010/webp"

	"github.com/nfnt/resize"

	"restapi/src/libs"
	"restapi/src/models"
	"restapi/src/pkg/config"
	"restapi/src/pkg/fileUtils"
	"restapi/src/pkg/imageUtils"
)

/*
func DownloadHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
*/
func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Add("Server", "SeusCode-Images/1.0")
	w.Header().Add("Seus", "Images/1.0")
	w.Header().Add("Cache-Control", "public, max-age=86400")

	if r.Method != "GET" {
		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", "GET")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	//Declaracion explicita de variables
	var requestURI *url.URL
	var err error

	//Parsear URL
	requestURI, err = url.ParseRequestURI(r.RequestURI)
	//Error al parsear o sin path?
	if err != nil || requestURI.Path == "" {
		fmt.Println("NO PATH IN REQUEST")
		http.ServeFile(w, r, "___ERROR.png")
		return
	}

	//Split del path
	spl := strings.Split(requestURI.Path, "/")

	// Example Path: /api/images/v1/download/15fc0db4-45c4-42c7-bd55-25fa7b0b9219_39f8cc344469e601b.webp
	// Example spl (split(path,"/")):
	// [0] => "", [1] => "api", [2] => "images", [3] => "v1", [4] => "download", [5] => "15fc0db4,-45c4-42c7-bd55-25fa7b0b9219_39f8cc344469e601b.webp"

	//previene que se inserten slashes extra en la url (spl debe siempre tener 6 elementos)
	if len(spl) != 6 {
		fmt.Println("WRONG PATH IN REQUEST")
		http.ServeFile(w, r, "___ERROR.png")
		return
	}

	//filename obtenido del ultimo elemento del path
	filename, err := url.QueryUnescape(spl[len(spl)-1])

	//No deberia pasar, pero por si ocurre un error
	if err != nil {
		fmt.Println("ERROR QueryUnescape(filename)")
		http.ServeFile(w, r, "___ERROR.png")
		return
	}

	//Hacemos un split del filename para obtener el id unico del archivo/registro
	spl = strings.Split(filename, "_")

	//si hay menos de 2 elementos entonces dar error (imagen)
	//no se cumple con el criterio: %id% + "_" + %original_name%
	if len(spl) < 2 {
		/*//insertar registro en log de base de datos
		libs.DB.Create(&models.SystemLog{
			Name: fmt.Sprintf("Service IMAGES (ERROR) (PID %d)", os.Getpid()),
			Desc: fmt.Sprintf("Error en UID_FILENAME SYNTAX"),
		})
		*/
		fmt.Println("ERROR len(split(filename))")
		http.ServeFile(w, r, "___ERROR.png")
		return
	}
	//usar spl[0] para el uid (unique id)
	uid := spl[0]

	//Reemplazar/quitar en filename el string del uid
	filename = strings.Replace(filename, uid+"_", "", 1)

	// construir el path de acceso al archivo real
	// Ej: images/a/b/c/abc-asdasd-as777_filename.ext
	// config.GetConfig().IMAGES_DIR = images
	// uid[0] = a
	// uid[1] = b
	// uid[2] = c
	// uid = abc-asdasd-as777
	// filename = filename.ext
	fPath, err := url.QueryUnescape(fmt.Sprintf("%s/%c/%c/%c/%s_%s", config.GetConfig().IMAGES_DIR, uid[0], uid[1], uid[2], uid, filename))

	//No deberia pasar, pero por si ocurre un error
	if err != nil {
		fmt.Println("ERROR QueryUnescape(fPath)")
		http.ServeFile(w, r, "___ERROR.png")
		return
	}

	//Si no existe o el nombre de archivo tiene menos de 3bytes
	if !fileUtils.FileExists(fPath) || len(filename) < 3 {
		fmt.Print("\n*----------------\n")
		fmt.Print("ERROR 404")
		fmt.Print("\n-----------------\n\n")
		//log.Printf("Redirecting %s to image404.jpg (HTTP Code %d).", fPath, nfrw.status)
		http.ServeFile(w, r, "___ERROR.png") //, http.StatusFound)
		return
	} else {
		fmt.Print("-----------------\n")
		fmt.Print("SIN ERROR")
		fmt.Print("\n-----------------\n\n")

		fSize, _ := fileUtils.FileSize(fPath)
		fmt.Println(fSize)
		result := models.Image{}

		//Hacer query para comprobar que el archivo existe en DB
		ret := libs.DB.Where(&models.Image{Uid: string(uid), Name: filename}).First(&result)

		//Si no tenemos resultados dar error
		if ret.RowsAffected < 1 {
			http.ServeFile(w, r, "___ERROR.png") //, http.StatusFound)
			return
		}

		//Si los datos obtenidos del path no coinciden con la db dar error
		//Esto tecnicamente no deberia ocurrir, quizas este de mas ya que se comprueban en el mismo query
		if result.Uid != uid || result.Name != filename {
			http.ServeFile(w, r, "___ERROR.png") //, http.StatusFound)
			return
		}

		//ret, _ := libs.DB.Select("id >= ?", 6).Rows() //.First(&result)
		//fmt.Print(ret.Statement)
		//libs.DB

		//fmt.Println(libs.DB.Where(&models.Image{Uid: string(uid)}).First(&result)) //.Updates(&models.Image{TotalTransferSize: uint(fSize)})
		//fmt.Printf("%+v", &result)
		fmt.Printf("Nombre del archivo: %s_%s\n", url.QueryEscape(result.Uid), url.QueryEscape(result.Name))

		//result.Uid
		fmt.Printf("UID %s\n", uid)

		//fmt.Printf("resultados %d\n", ret.RowsAffected)

		//libs.DB.Model(&models.Image{})

		/*var imgfile io.Reader //ReadCloser
		var err error
		/**/

		var imgfile *os.File
		imgfile, err = os.Open(fPath)
		if err != nil {
			fmt.Print(err.Error())
		}
		//defer imgfile.Close()

		fmt.Print("*******************************\n")
		fmt.Print("*******************************\n")

		//filetype := "image/png"
		filetype := fileUtils.GuessImageMimeTypes(imgfile)
		imgfile.Seek(0, 0)

		var img image.Image

		fmt.Println(filetype)
		fmt.Println(imgfile)

		switch filetype {
		case "image/jpeg", "image/jpg":
			// decode jpeg into image.Image
			img, err = jpeg.Decode(imgfile)
			if err != nil {
				fmt.Println(err)
				//log.Fatal(err)
			}

		case "image/gif":
			// decode jpeg into image.Image
			img, err = gif.Decode(imgfile)
			if err != nil {
				fmt.Println(err)
				//log.Fatal(err)
			}

		case "image/png":
			// decode jpeg into image.Image
			img, err = png.Decode(imgfile)
			//fmt.Println(img)
			fmt.Println(err)
			if err != nil {
				fmt.Println(err)
				//log.Fatal(err)
			}

		case "image/webp":
			// decode jpeg into image.Image
			img, err = webp.Decode(imgfile)
			if err != nil {
				fmt.Println(err)
				//log.Fatal(err)
			}

			fmt.Println(filetype)
		case "image/bmp":
			// decode jpeg into image.Image
			img, err = bmp.Decode(imgfile)
			if err != nil {
				fmt.Println(err)
				//log.Fatal(err)
			}

		default:
			fmt.Println("unknown file type uploaded")
			http.ServeFile(w, r, "___ERROR.png")
			return
		}

		origBounds := img.Bounds()
		origWidth := uint(origBounds.Dx())
		origHeight := uint(origBounds.Dy())
		/*
			// Nearest-neighbor interpolation
			NearestNeighbor InterpolationFunction = iota
			// Bilinear interpolation
			Bilinear
			// Bicubic interpolation (with cubic hermite spline)
			Bicubic
			// Mitchell-Netravali interpolation
			MitchellNetravali
			// Lanczos interpolation (a=2)
			Lanczos2
			// Lanczos interpolation (a=3)
			Lanczos3
		*/
		if origWidth > 1920 || origHeight > 1080 {
			imgFHD := resize.Thumbnail(1920, 1080, img, resize.NearestNeighbor)
			imageUtils.ImageSave(fPath+".normal1920."+strings.Replace(filetype, "image/", "", -1), imgFHD)

		}

		if origWidth > 1024 || origHeight > 768 {
			imgNormal := resize.Thumbnail(1024, 768, img, resize.NearestNeighbor)
			imageUtils.ImageSave(fPath+".normal1024."+strings.Replace(filetype, "image/", "", -1), imgNormal)

		}

		if origWidth > 480 || origHeight > 360 {
			imgThumb := resize.Thumbnail(480, 360, img, resize.NearestNeighbor)
			imageUtils.ImageSave(fPath+".thumb480."+strings.Replace(filetype, "image/", "", -1), imgThumb)

		}

		if origWidth > 320 || origHeight > 240 {
			imgThumb := resize.Thumbnail(320, 240, img, resize.NearestNeighbor)
			imageUtils.ImageSave(fPath+".thumb320."+strings.Replace(filetype, "image/", "", -1), imgThumb)

		}

		imgfile.Close()

		fmt.Print("|||||||||||||||||||||||||||||||\n")

		http.ServeFile(w, r, fPath) //"relative/path/to/favicon.ico")
	}

	/*
		nfrw := &NotFoundRedirectRespWr{ResponseWriter: w}

		fmt.Print(r)
		fmt.Print("\n====================================================================\n\n")

		fmt.Print("-----------------\n")
		fmt.Print(r.RequestURI)
		fmt.Print("\n-----------------\n\n")

		h.ServeHTTP(nfrw, r)

		if nfrw.status == 404 || len(strings.Replace(r.RequestURI, "/"+config.GetConfig().IMAGES_DIR+"/", "", -1)) < 8 {
			fmt.Print("-----------------\n")
			fmt.Print("ERROR 404")
			fmt.Print("\n-----------------\n\n")
			//log.Printf("Redirecting %s to image404.jpg (HTTP Code %d).", r.RequestURI, nfrw.status)
			http.Redirect(w, r, "/images/_/_/_/___ERROR.jpg", http.StatusFound)
		} else {

			fmt.Print("-----------------\n")
			fmt.Print("SIN ERROR")
			fmt.Print("\n-----------------\n\n")
		}
	*/

	/*
		if nfrw.status == 404 {
			log.Printf("Redirecting %s to index.html.", r.RequestURI)
			http.Redirect(w, r, "/index.html", http.StatusFound)
		}
	*/
}

//}
