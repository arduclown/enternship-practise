package main

import (
	"fmt"
	"net/http"

	"github.com/arduclown/enternship-practise/utils"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/student", student)
	http.HandleFunc("/students", students)
	http.HandleFunc("/students/grade", grade)
	http.HandleFunc("/student/add", addStudent)
	http.HandleFunc("/students/grade-stats", gradeStats)

	http.HandleFunc("/students/age-stats", studentAge)
	utils.LoadStudentFromFile(&baseStud)

	if err := utils.InitDB(); err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		return
	}
	defer utils.CloseDB()

	http.ListenAndServe(":8080", nil)
}
