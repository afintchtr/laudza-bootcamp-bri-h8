package main

import (
	"fmt"
	"golang-days/day05/funcs"
	"strings"
)

func happyMonke(a string, b string) bool {
	if (strings.ToLower(a) == strings.ToLower(b)) {
		return false
	}
	return true
}

func scaryRoster(time uint, isCrowed bool) bool {
	if (isCrowed) {
		if (time <= 3 || time == 12 || time == 24) {
			return true
		} 
	}
	return false
}

func firsTask () {
	a := "sEnyUm"
	b := "cEmbeRut"
	if (happyMonke(a, b)) {
		fmt.Println("all good")
	} else {
		fmt.Println("in trouble")
	}

	var time uint = 2
	isCrowed := false
	if (scaryRoster(time, isCrowed)) {
		fmt.Println("u r in danger")
	} else {
		fmt.Println("all good")
	}
}

func main () {
	funcs.AnArray()
}