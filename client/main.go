package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	log.Println("----IP Test----")
	privateIP, publicIP, handlerAddress := getNetInfo()
	log.Println("Handler Server Address:", handlerAddress)
	log.Println("Private IP:", privateIP)
	log.Println("Public IP:", publicIP)

}

func getNetInfo() (privateIP, publicIP, handlerServerAddr string) {
	// 1. Get the home directory
	homeDir, err := os.UserHomeDir()
	panicOnErr(err)
	// 2. Get the ip info file
	addressFile, err := os.ReadFile(homeDir + "/info/network.json")
	panicOnErr(err)
	// 3. Extract the data
	addressFileData := map[string]string{}
	err = json.Unmarshal(addressFile, &addressFileData)
	panicOnErr(err)
	publicIP = addressFileData["pub_ip"]
	privateIP = addressFileData["pr_ip"]
	handlerServerAddr = addressFileData["handlerAddress"]
	return
}

func panicOnErr(err error) {
	if err != nil {
		log.Println("Error:", err.Error())
		panic(err)
	}
}
