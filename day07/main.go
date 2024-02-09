package main

import (
	"golang-days/day07/database"
)

// "golang-days/day07/funcs"

func main() {
	// funcs.PrintAllUkerFuncs()
	// api.MainApi()

	db := database.ConnectDB()
	database.CreateEmployee(db)

	defer db.Close()
}