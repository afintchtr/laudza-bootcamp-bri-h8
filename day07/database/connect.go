package database

import (
	"database/sql"
	"fmt"
	"golang-days/day07/env"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	err error
)

func ConnectDB() (*sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", env.Host, env.Port, env.User, env.Password, env.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	return db
}

// type Employee struct {
// 	ID        int
// 	Full_name string
// 	Email     string
// 	Age       int
// 	Division  string
// }

// func CreateEmployee(connection *sql.DB) {
// 	db := connection
// 	var employee = Employee{}

// 	sqlStatement := `
// 	INSERT INTO employee (full_name, email, age, division)
// 	VALUES ($1, $2, $3, $4)
// 	Returning *
// 	`

// 	err = db.QueryRow(sqlStatement, "Laudza", "laudza@test.com", 23, "APP").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("New Employee Data: %+v\n", employee)
// }