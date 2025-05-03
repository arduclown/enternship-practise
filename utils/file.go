package utils

import (
	"encoding/json"
	"os"
)

// check errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadStudentFromFile(students *[]Student) {

	f, err := os.Open("./students.txt")
	check(err)
	defer f.Close()

	var newStudents []Student
	decoder := json.NewDecoder(f)
	decoder.Decode(&newStudents)

	*students = newStudents
}
