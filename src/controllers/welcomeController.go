/*

 * Original Idea: Sergio Baudcracco. First code: Genaro Civilotti
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
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API REST\n")

	fmt.Print("-----------------\n")
	fmt.Print(r.RequestURI)
	fmt.Print("\n-----------------\n\n")
	w.Header().Add("Server", "SeusCode-Images/1.0")
	w.Header().Add("Seus", "Images/1.0")
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Cache-Control", "public, max-age=86400")

	w.WriteHeader(200)
	w.Write([]byte("\n\nNovedades:\nMejora del servicio, redimensiona imagenes subidas (actualmente se hace en downloadController)\nAgregados modelos images, logs y users\nAgregada compatibilidad TLS/SSL\n\nTODO List:\nRedimensionar imagenes en uploadController en lugar de downloadController (mover y adaptar rutina)\nRedimensionar imagenes segun tipo de usuario (o su nivel de membresia)\nAgregar endpoint para obtener el thumbnail de la imagen (actualmente se genera el thumb en downloadController y se almacena en disco)\nAutomatizar borradod e imagenes que no cumplan con requisitos minimos de visitas/actividad (hacer configurable en config.json) siempre y cuando el nivel/membresia del usuario no sobreeescriba esta politica\nCrear endpoint que devuelva los stats de las imagenes segun informacion de la DB, por imagen y rankings generales -a determinar-\nPoner un limite de transferencia por imagen segun nivel o tipo de usuario propietario de la imagen (al agotar el limite, dar imagen de error)\nAgregar el archivo favicon.ico (aunque sea en blanco)"))

}
