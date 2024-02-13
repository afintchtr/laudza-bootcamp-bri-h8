package main

import (
	"golang-days/day10-jwt/database"
	"golang-days/day10-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8081")
}