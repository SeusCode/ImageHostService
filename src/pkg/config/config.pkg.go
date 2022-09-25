package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ConfigStruct struct which contains a configurations
type ConfigStruct struct {
	SERVER_PORT      int    `json:"SERVER_PORT"`
	IMAGES_DIR       string `json:"IMAGES_DIR"`
	MAX_FILE_SIZE    int64  `json:"MAX_FILE_SIZE"`
	MAX_FILES_UPLOAD int    `json:"MAX_FILES_UPLOAD"`
	MAX_THREADS      int    `json:"MAX_THREADS"`
	IMAGES_FORM      string `json:"IMAGES_FORM"`
}

var serverConfig ConfigStruct
var initialized bool

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

func GetConfig() ConfigStruct {
	if initialized == false {
		Load()
	}
	return serverConfig
}

func GetMaxUploadSize() int64 {
	if initialized == false {
		Load()
	}
	return int64(serverConfig.MAX_FILES_UPLOAD) * serverConfig.MAX_FILE_SIZE
}
