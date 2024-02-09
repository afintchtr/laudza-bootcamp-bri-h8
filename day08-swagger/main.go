package main

import (
	"golang-days/day08-swagger/database"
	"golang-days/day08-swagger/routers"
)

func main() {
	var PORT = ":3030"

	database.StartDB()
	routers.StartServer().Run(PORT)
}