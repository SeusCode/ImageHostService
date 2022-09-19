package routes

import (
	"net/http"
	"restapi/src/controllers"
)

func Setup() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/upload", controllers.Upload)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
}