package routes

import (
	"net/http"
	"restapi/src/controllers"
)

func Setup() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/api/images/v1/upload", controllers.Upload)
	//http.HandleFunc("/api/images/v1/token", controllers.Token)

	//http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	http.HandleFunc("/api/images/v1/download/", controllers.DownloadHandler)
	//http.HandleFunc("/api/images/v1/download/", controllers.DownloadHandler(http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))))
}
