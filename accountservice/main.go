package main

import (
	"fmt"
	"github.com/credondocr/goblog/accountservice/services"
)

var appName = "accountservice"


func main() {
	fmt.Printf("Starting %v\n", appName)
	services.StartWebServer("6767")
}