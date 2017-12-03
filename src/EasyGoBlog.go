package main

import (
	"bServer"
	"databaseTools"
	"fmt"
)

var myServer bServer.WebServer

func init() {
	//myServer.InitTemplates()
}

func main() {
	//myServer.StartServer()

	var mysqlTool databaseTools.MysqlTool
	err := mysqlTool.ConnectDatabase("test", "root", "123456", "localhost", "3306")
	CheckErr(err)

	err = mysqlTool.AddUser("5", "test")
	CheckErr(err)

	err = mysqlTool.DeleteUser("3")
	CheckErr(err)

	mysqlTool.CloseConnect()

}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
