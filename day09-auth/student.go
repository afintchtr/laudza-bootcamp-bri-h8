package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func init() {
	students = append(students, &Student{Id: "1", Name: "Laudza", Grade: 7})
	students = append(students, &Student{Id: "2", Name: "Raka", Grade: 9})
	students = append(students, &Student{Id: "3", Name: "Azwar", Grade: 12})
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}
	return nil
}