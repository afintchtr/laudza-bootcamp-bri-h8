package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// var (
// 	db *sql.DB
// 	err error
// )

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func CreateEmployee(connection *sql.DB) {
	db := connection
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employee (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	err = db.QueryRow(sqlStatement, "Rosa", "rosa@test.com", 28, "ISC").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data: %+v\n", employee)
}