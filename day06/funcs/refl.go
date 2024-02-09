package funcs

import (
	"fmt"
	"reflect"
)

type Organization struct {
	orgName string
	address string
}

type Employee struct {
	nik      string
	name     string
	age      int
	division string
	org      Organization
}

func PracticeReflectInStruct() {
	employees := []Employee{
		{nik: "3202180708090007", name: "Isnanda Muhammad Zain", age: 35, division: "DDB", org: Organization{orgName: "BRI", address: "Bendungan Hilir"}},
		{nik: "3202180708090016", name: "Muhammad Azwar Rasyid", age: 28, division: "APP", org: Organization{orgName: "BCA", address: "Bendungan Beler"}},
		{nik: "3202180708090123", name: "Dionysius Raka Bening", age: 40, division: "EDM", org: Organization{orgName: "BSI", address: "Bendungan Tolil"}},
	}
	for _, employee := range employees {
		fmt.Println(reflect.ValueOf(employee), reflect.TypeOf(employee))
	}

	
	fmt.Println(findOldest(employees))
}

func findOldest(employees []Employee) Employee {
	oldestAge := -1
	var oldestEmployee Employee
	for _, employee := range employees {
		if employee.age > oldestAge {
			oldestAge = employee.age
			oldestEmployee = employee
		}
	}
	return oldestEmployee
}