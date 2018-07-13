package main

import (
	"fmt"
	"github.com/credondocr/goblog/accountservice/dbclient"
	"github.com/credondocr/goblog/accountservice/services"
)

var appName = "accountservice"


func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()                 // NEW
	services.StartWebServer("6767")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	services.DBClient = &dbclient.BoltClient{}
	services.DBClient.OpenBoltDb()
	services.DBClient.Seed()
}