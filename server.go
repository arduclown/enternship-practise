// маршруты
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arduclown/enternship-practise/utils"
)

var baseStud = []utils.Student{
	{Name: "Ann", Age: 16, Grade: 3.85},
	{Name: "Anton", Age: 16, Grade: 4},
	{Name: "Bob", Age: 17, Grade: 5},
	{Name: "Sarah", Age: 16, Grade: 4.9},
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

// вывод студента по ИМЕНИ
func student(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	w.Header().Set("Content-Type", "application/json")
	for _, s := range baseStud {
		if s.Name == name {
			json.NewEncoder(w).Encode(s)
			return
		}
	}
	if err := json.NewEncoder(w).Encode(map[string]string{"error": "Student not found"}); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}

}

// вывод списка ВСЕХ студентов
func students(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(baseStud); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}

// вывод списка студентов, у которых ср. балл не ниже минимума
func grade(w http.ResponseWriter, req *http.Request) {
	grade := req.URL.Query().Get("min")
	w.Header().Set("Content-Type", "application/json")
	g, err := strconv.ParseFloat(grade, 64)
	filtred := []utils.Student{}
	if err != nil {
		http.Error(w, "Invalid grade parameter", http.StatusBadRequest)
		return
	}
	for _, s := range baseStud {
		if s.Grade >= g {
			filtred = append(filtred, s)
		}
	}
	json.NewEncoder(w).Encode(filtred)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/student", student)
	http.HandleFunc("/students", students)
	http.HandleFunc("/students/grade", grade)

	utils.LoadStudentFromFile(&baseStud)

	http.ListenAndServe(":8080", nil)
}
