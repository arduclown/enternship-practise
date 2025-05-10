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
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
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
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	base, _ := utils.GetStudents()
	if err := json.NewEncoder(w).Encode(base); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}

// вывод списка студентов, у которых ср. балл не ниже минимума
func grade(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
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

func addStudent(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var student utils.Student
	if err := json.NewDecoder(req.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	baseStud = append(baseStud, student)
	utils.SaveToFile(baseStud)

	if utils.InsertStudent(student) != nil {
		http.Error(w, "Failes to save", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Student added"}); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/student", student)
	http.HandleFunc("/students", students)
	http.HandleFunc("/students/grade", grade)
	http.HandleFunc("/student/add", addStudent)

	utils.LoadStudentFromFile(&baseStud)

	if err := utils.InitDB(); err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		return
	}

	http.ListenAndServe(":8080", nil)
}
