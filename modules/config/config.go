package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigStruct struct {
	Server ConfigServer
	MSF    MSFServer
}

type ConfigServer struct {
	Port int `json:"port"`
}

type MSFServer struct {
	Location string `json:"location"`
}

const configFile string = "mycloud.config"

var DebugMode bool = false
var config ConfigStruct = ConfigStruct{
	Server: ConfigServer{
		Port: 8080,
	},
	MSF: MSFServer{
		Location: "./",
	},
}

func init() {
	checkFile()
	log.Println("Load config file...")
	config = loadConfiguration()
}

func GetServerPort() int {
	return config.Server.Port
}

func GetRootLocation() string {
	return config.MSF.Location
}

func loadConfiguration() ConfigStruct {
	var result ConfigStruct
	if fileObj, fileErr := ioutil.ReadFile(configFile); fileErr == nil {
		if jsonErr := json.Unmarshal(fileObj, &result); jsonErr != nil {
			log.Printf("File %s has an error! Cannot continue with the server, please check the syntax!", configFile)
			log.Panicln(jsonErr.Error())
		}
	}

	return result
}

func checkFile() {
	fileObj, err := os.OpenFile(configFile, os.O_WRONLY, 600)

	defer fileObj.Close()

	if err != nil {
		newFileObj, _ := os.Create(configFile)

		defer newFileObj.Close()

		jsonBytes, jsonBytesErr := json.MarshalIndent(config, "", "\t")
		if jsonBytesErr == nil {
			newFileObj.Write(jsonBytes)
		}
	}
}
