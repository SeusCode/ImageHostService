package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi/src/pkg/config"
	"restapi/src/routes"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("ImageHoster v1")
	fmt.Println("by SEUS\n\n")

	//Esto no es necesario, Init() es llamado desde las funciones del paquete config en caso de que no este inicializado (config cargada)
	//config.Load()

	if runtime.NumCPU() >= config.GetThreads() {
		hilos := config.GetThreads()

		fmt.Println("Cantidad de procesadores:", runtime.NumCPU())
		fmt.Printf("Habilitando %d hilos para el proceso\n", hilos)

		runtime.GOMAXPROCS(hilos)

	}

	wg.Add(2)

	go func() {
		//fmt.Println("Goroutine 1")
		//TODO: Do schedule tasks in a loop
		wg.Done()
	}()

	go func() {
		fmt.Println("Server Goroutine start")

		routes.Setup()
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.GetPort()), nil)

		if err != nil {
			log.Fatal(err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Exit")

}
