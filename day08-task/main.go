package main

import (
	"golang-days/day08-task/database"
	"golang-days/day08-task/routers"
)

func main() {
	var PORT = ":7070"

	database.StartDB()
	routers.StartServer().Run(PORT)
}