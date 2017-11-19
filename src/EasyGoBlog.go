package main

import (
	"bServer"
)

var myServer bServer.WebServer

func init() {
	myServer.InitTemplates()
}

func main() {
	myServer.StartServer()
}

