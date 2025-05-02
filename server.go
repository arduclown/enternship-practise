// маршруты
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseStud = []Student{
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
	json.NewEncoder(w).Encode(map[string]string{"error": "Student not found"})

}

// вывод списка ВСЕХ студентов
func students(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(baseStud); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/student", student)
	http.HandleFunc("/students", students)

	http.ListenAndServe(":8080", nil)

	// nums := []int{1, 2, 3, 4, 5}
	// ReverseArray(nums)
	// fmt.Println(nums)

}
