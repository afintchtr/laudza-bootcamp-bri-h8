package main

import (
	"golang-days/day08-gin/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}