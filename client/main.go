package main

import (
	"context"
	"ec2-grpc-ip-test/client/IPTestService"
	"encoding/json"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Println("----IP Test----")
	privateIP, publicIP, handlerAddress := getNetInfo()
	log.Println("Handler Server Address:", handlerAddress)
	log.Println("Private IP:", privateIP)
	log.Println("Public IP:", publicIP)

	// Call the server
	conn, err := grpc.Dial(handlerAddress+":2500", grpc.WithTransportCredentials(insecure.NewCredentials()))
	panicOnErr(err)
	defer func() {
		conn.Close()
	}()
	client := IPTestService.NewIP_TestClient(conn)

	res, err := client.GetIP(context.Background(), &IPTestService.Req{
		Msg: "Just for testing purposes",
	})

	log.Println("Response IP:", res.Ip)
	log.Println("Response Address:", res.Address)
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
