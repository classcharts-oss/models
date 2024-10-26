package main

import (
	"fmt"

	"github.com/classcharts-oss/models/student"
)

func main() {
	student := student.NewUser(1, "Jeremy Kyle", "https://placehold.co/360")

	Printlnf("ID: %d, Name: %s, First Name: %s, Last Name: %s", student.Id, student.Name, student.FirstName, student.LastName)
}

func Printlnf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}
