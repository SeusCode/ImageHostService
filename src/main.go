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

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restapi/src/libs"
	"restapi/src/models"
	"restapi/src/pkg/config"
	"restapi/src/pkg/fileUtils"
	"restapi/src/routes"
	"runtime"
	"sync"
	"syscall"
	"time"

	task "github.com/SERBice/gotaskscheduler"
)

var wg sync.WaitGroup

func main() {

	fmt.Printf("%-55s", fmt.Sprintf("%55s", "------------------\n"))
	fmt.Printf("%-55s", fmt.Sprintf("%55s", "- ImageHoster v1 -\n"))
	fmt.Printf("%-55s", fmt.Sprintf("%55s", "-    By SEUS     -\n"))
	fmt.Printf("%-55s", fmt.Sprintf("%55s", "------------------\n"))

	fmt.Print("\n\nAPI v1.0.1\n\n")
	/*
		fmt.Println("1----------------------------")
		basepath, _ := os.Getwd() //.Executable()

		//basepath = filepath.ToSlash(basepath)
		//basepath = path.Dir(basepath)
		fmt.Println(basepath)
		//midir, mifile := path.Split(basepath)
		//fmt.Println(midir)
		//fmt.Println(mifile)
		//fmt.Println(filepath.ToSlash(basepath))
		fmt.Println(path.Join(basepath, "ssl.ca"))
		fmt.Println(fileUtils.FileExists(path.Join(basepath, "ssl.ca")))
		os.Exit(0)/**/
	/*
		fmt.Println("2----------------------------")
		basepath = path.Dir(basepath)
		fmt.Println(basepath)
		midir, mifile = path.Split(basepath)
		fmt.Println(midir)
		fmt.Println(mifile)
		//fmt.Println(filepath.ToSlash(basepath))
		os.Exit(0)
	*/

	dbConfig := libs.DbConfig{
		Host:     config.GetConfig().DBHOST,
		Port:     config.GetConfig().DBPORT,
		Database: config.GetConfig().DBNAME,
		User:     config.GetConfig().DBUSER,
		Password: config.GetConfig().DBPASS,
		Charset:  config.GetConfig().DBCHAR,
	}

	libs.DB = dbConfig.InitMysqlDB()

	libs.DB.Create(&models.SystemLog{
		Name: fmt.Sprintf("Service IMAGES (BOOT) (PID %d)", os.Getpid()),
	})

	//Esto no es necesario, Init() es llamado desde las funciones del paquete config en caso de que no este inicializado (config cargada)
	//config.Load()

	if runtime.NumCPU() >= config.GetConfig().MAX_THREADS {
		hilos := config.GetConfig().MAX_THREADS

		fmt.Println("Cantidad de procesadores:", runtime.NumCPU())
		fmt.Printf("Habilitando %d hilos para el proceso\n", hilos)

		runtime.GOMAXPROCS(hilos)

		libs.DB.Create(&models.SystemLog{
			Name: fmt.Sprintf("Service IMAGES (THREADS: %d) (PID %d)", config.GetConfig().MAX_THREADS, os.Getpid()),
		})

	}

	//ret :=
	libs.DB.Create(&models.SystemLog{
		Name: fmt.Sprintf("Service IMAGES (BOOT OK) (PID %d)", os.Getpid()),
	})
	//fmt.Println(ret.Error)        // devuelve el error
	//fmt.Println(ret.RowsAffected) // devuelve la cantidad de registros creados

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	wg.Add(3)

	go func() {
		<-c

		// Run Cleanup
		fmt.Println("Please wait...")
		//Realizar tareas de guardado
		task.StopScheduler(false)

		//esperar 1500ms
		time.Sleep(2000 * time.Millisecond)

		os.Exit(1)
	}()

	go func() {
		//fmt.Println("Goroutine 1")
		//TODO: Do schedule tasks in a loop

		/*
			fmt.Println("Preparando tareas recurrentes...")
			task.SetTasksLimit(4000)
			task.AddTask("", 1, func() {
				return
			}, true)
			task.AddTask("", 1, func() {
				for {
					time.Sleep(1 * time.Second)
					fmt.Print(":")
				}
			}, true)
			task.AddTask("", 1, func() {
				return
			}, true)
			task.AddTask("", 1, func() {
				return
			}, true)
			task.AddTask("", 1, func() { fmt.Print(".") }, false)
			fmt.Println("Se han programado ", task.CountTasks(), " tareas.")
			task.StartScheduler()
			fmt.Println("Despues de 0 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.1 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.2 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.3 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.4 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.5 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.6 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.7 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.8 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 0.9 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(111 * time.Millisecond)
			fmt.Println("Despues de 1 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(2 * time.Second)
			fmt.Println("Despues de 3 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(2 * time.Second)
			fmt.Println("Despues de 5 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(6 * time.Second)
			fmt.Println("Despues de 11 segundos quedan", task.CountTasks(), "tareas.")
			time.Sleep(20 * time.Second)
			fmt.Println("Despues de 31 segundos quedan", task.CountTasks(), "tareas.")
		*/
		wg.Done()
	}()

	go func() {

		libs.DB.Create(&models.SystemLog{
			Name: fmt.Sprintf("Service IMAGES (SERVER START) (PID %d)", os.Getpid()),
		})

		fmt.Println("Server start")

		routes.Setup()

		var err error

		if config.GetConfig().SERVER_SSL == true && (!fileUtils.FileExists("ssl.cert") || !fileUtils.FileExists("ssl.key")) {
			n := 0
			for n < 3 {
				fmt.Println("SSL Server is Enabled but Cert File and/or Key File do not exist.")
				n++
			}
			fmt.Println("Please Check ssl.cert and ssl.key, or disable SSL Server.")
		}

		if config.GetConfig().SERVER_SSL == true && fileUtils.FileExists("ssl.cert") && fileUtils.FileExists("ssl.key") {
			err = http.ListenAndServeTLS(fmt.Sprintf(":%d", config.GetConfig().SERVER_PORT), "ssl.cert", "ssl.key", nil)
		} else {
			err = http.ListenAndServe(fmt.Sprintf(":%d", config.GetConfig().SERVER_PORT), nil)
		}

		if err != nil {
			log.Fatal(err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
}
