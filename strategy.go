package main

import "fmt"

type IDBconnection interface {
	Connect()
}

type SqlConnection struct {
	connectionString string
}

func (con SqlConnection) Connect() {
	fmt.Println("SQL " + con.connectionString)
}

type OracleConnection struct {
	connectionString string
}

func (con OracleConnection) Connect() {
	fmt.Println("Oracle " + con.connectionString)
}

type DBConnection struct {
	db IDBconnection
}

func (con DBConnection) DBConnect() {
	con.db.Connect()
}

func main() {
	sqlConnection := SqlConnection{
		connectionString: "Connection is connected!",
	}
	con1 := DBConnection{
		db: sqlConnection,
	}
	con1.db.Connect()
	oracleConnection := OracleConnection{
		connectionString: "Connection is connected!",
	}
	con2 := DBConnection{
		db: oracleConnection,
	}
	con2.db.Connect()
}
