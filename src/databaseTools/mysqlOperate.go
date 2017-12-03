package databaseTools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//mysql database operate type
//time:2017-12-03
type MysqlTool struct {
	databaseName     string
	databaseUser     string
	databasePassword string

	databaseLocation string
	databasePort     string

	database *sql.DB
}

const (
	constMaxConnection = 1000 // Max connection numbers
)

//basic connect to mysql
func (mysqlTool *MysqlTool) ConnectDatabase(dbName, dbUser, dbPassword, dbLocation, dbPort string) (err error) {
	err = nil

	mysqlTool.databaseName = dbName
	mysqlTool.databaseUser = dbUser
	mysqlTool.databasePassword = dbPassword
	mysqlTool.databaseLocation = dbLocation
	mysqlTool.databasePort = dbPort

	defer func() {
		if dErr := recover(); dErr != nil {
			fmt.Println(dErr, "")
		}
	}()

	//check
	if mysqlTool.databaseName == "" || mysqlTool.databaseUser == "" || mysqlTool.databaseLocation == "" || mysqlTool.databasePort == "" {
		panic("Exception: reference string is empty!")
		return
	}
	connectionOrder := mysqlTool.databaseUser + ":" + mysqlTool.databasePassword + "@tcp(" + mysqlTool.databaseLocation + ":" + mysqlTool.databasePort + ")/" + mysqlTool.databaseName + "?charset=utf8"
	//fmt.Println("Order: ", connectionOrder)
	mysqlTool.database, err = sql.Open("mysql", connectionOrder)
	mysqlTool.database.SetMaxOpenConns(constMaxConnection)
	err = mysqlTool.database.Ping()
	return
}

//Close connection
func (mysqlTool *MysqlTool) CloseConnect() {
	if mysqlTool.database != nil {
		err := mysqlTool.database.Close()
		if err != nil {
			fmt.Println("Exception:", err)
		}
	}
}

func (mysqlTool *MysqlTool) AddUser(id, name string) error {
	stmt, err := mysqlTool.database.Prepare("INSERT user (id, name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	theId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(theId)

	return nil
}

func (mysqlTool *MysqlTool) DeleteUser(id string) error {
	stmt, err := mysqlTool.database.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(num)

	return nil
}

func (mysqlTool *MysqlTool) UpdateUser(id, name string) error {



	return nil
}
