package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// jsonCfg struct which contains a configurations
type jsonCfg struct {
	SERVER_PORT      int    `json:"SERVER_PORT"`
	IMAGES_DIR       string `json:"IMAGES_DIR"`
	MAX_FILE_SIZE    int    `json:"MAX_FILE_SIZE"`
	MAX_FILES_UPLOAD int    `json:"MAX_FILES_UPLOAD"`
	MAX_THREADS      int    `json:"MAX_THREADS"`
	IMAGES_FORM      string `json:"IMAGES_FORM"`
}

var serverConfig jsonCfg
var initialized bool

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func Load() {

	fmt.Println("Loading config.json...")
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("Successfully opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'serverConfig' which we defined above
	json.Unmarshal(byteValue, &serverConfig)

	if serverConfig.SERVER_PORT == 0 || serverConfig.MAX_FILE_SIZE == 0 || serverConfig.MAX_FILES_UPLOAD == 0 || serverConfig.IMAGES_DIR == "" {
		log.Fatal("Invalid config file")
		return
	}

	fmt.Println("Successfully loaded config.json")

	initialized = true
}

func GetThreads() int {
	if initialized == false {
		Load()
	}
	return serverConfig.MAX_THREADS
}

func GetPort() int {
	if initialized == false {
		Load()
	}
	return serverConfig.SERVER_PORT
}

func GetImagesDir() string {
	if initialized == false {
		Load()
	}
	return serverConfig.IMAGES_DIR
}

func GetImagesForm() string {
	if initialized == false {
		Load()
	}
	return serverConfig.IMAGES_FORM
}

func GetMaxFileSize() int {
	if initialized == false {
		Load()
	}
	return serverConfig.MAX_FILE_SIZE
}

func GetMaxFilesUpload() int {
	if initialized == false {
		Load()
	}
	return serverConfig.MAX_FILES_UPLOAD
}

func GetMaxUploadSize() int {
	if initialized == false {
		Load()
	}
	return GetMaxFilesUpload() * GetMaxFileSize()
}
