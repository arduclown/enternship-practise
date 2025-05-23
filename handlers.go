package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/arduclown/enternship-practise/utils"
)

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

	var s utils.Student
	row := utils.DB.QueryRow("SELECT name, age, grade FROM students WHERE name = ?", name)
	err := row.Scan(&s.Name, &s.Age, &s.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(map[string]string{"error": "Student not found"})
			return
		}
		http.Error(w, "Failed to get student", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(s); err != nil {
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

	base, err := utils.GetStudents()
	if err != nil {
		http.Error(w, "Failed to get students", http.StatusInternalServerError)
		return
	}

	if len(base) == 0 {
		w.Header().Set("Content Type", "application/json")
		json.NewEncoder(w).Encode(ResponseAge{Students: []utils.Student{}, AverAge: 0})
		return
	}

	ages := make(chan int, len(base))
	var wg sync.WaitGroup

	for i := range base {
		wg.Add(1)
		go func(student utils.Student) {
			defer wg.Done()
			ages <- student.Age
		}(base[i])
	}

	go func() {
		wg.Wait()
		close(ages)
	}()

	var totalAge int
	count := 0

	for age := range ages {
		totalAge += age
		count++
	}

	averageAge := float64(totalAge) / float64(count)

	response := ResponseAge{
		Students: base,
		AverAge:  averageAge,
	}
	if err := json.NewEncoder(w).Encode(base); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func studentAge(w http.ResponseWriter, req *http.Request) {
	base, err := utils.GetStudents()
	if err != nil {
		http.Error(w, "Failed to get students", http.StatusInternalServerError)
		return
	}

	if len(base) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponseAge{Students: []utils.Student{}, AverAge: 0, MaxAge: 0})
		return
	}

	ageChan := make(chan int, len(base)) // для всех возрастов
	avgChan := make(chan float64, 1)     // средний возраст
	maxChan := make(chan int, 1)         // максимальный возраст
	var wg sync.WaitGroup

	wg.Add(1) //горутина для отправки возрастов
	go func() {
		defer wg.Done()
		for _, student := range base {
			ageChan <- student.Age
		}
		close(ageChan)
	}()

	var ages []int
	for age := range ageChan {
		ages = append(ages, age)
	}

	wg.Add(1) //горутина для вычисления среднего возраста
	go func(ages []int) {
		defer wg.Done()
		var totalAge int
		cnt := 0
		for _, age := range ages {
			totalAge += age
			cnt++
		}
		avg := float64(totalAge) / float64(cnt)
		avgChan <- avg
	}(ages)

	wg.Add(1)
	go func(ages []int) {
		defer wg.Done()
		maxAge := 0
		for _, age := range ages {
			if age > maxAge {
				maxAge = age
			}
		}
		maxChan <- maxAge
	}(ages)

	wg.Wait()

	averageAge := <-avgChan
	maxAge := <-maxChan

	response := ResponseAge{
		Students: base,
		AverAge:  averageAge,
		MaxAge:   maxAge,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
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

func gradeStats(w http.ResponseWriter, req *http.Request) {
	base, err := utils.GetStudents()
	if err != nil {
		http.Error(w, "Failed to get students", http.StatusInternalServerError)
		return
	}

	if len(base) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponseGrade{Students: []utils.Student{}, AverGrade: 0, MaxGrade: 0})
		return
	}

	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowes", http.StatusMethodNotAllowed)
		return
	}

	gradeChan := make(chan float64, len(base))
	avgChan := make(chan float64, 1)
	maxChan := make(chan float64, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, student := range base {
			gradeChan <- student.Grade
		}
		close(gradeChan)
	}()

	var grades []float64
	for grade := range gradeChan {
		grades = append(grades, grade)
	}

	wg.Add(1)
	go func(grades []float64) {
		defer wg.Done()
		var totalGrade float64
		cnt := 0
		for _, grade := range grades {
			totalGrade += grade
			cnt++
		}
		avg := float64(totalGrade) / float64(cnt)
		avgChan <- avg
	}(grades)

	wg.Add(1)
	go func(grades []float64) {
		defer wg.Done()
		maxGrade := 0.
		for _, grade := range grades {
			if grade > maxGrade {
				maxGrade = grade
			}
		}
		maxChan <- float64(maxGrade)
	}(grades)

	wg.Wait()

	averageGrade := <-avgChan
	maxGrade := <-maxChan

	response := ResponseGrade{
		Students:  base,
		AverGrade: averageGrade,
		MaxGrade:  maxGrade,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}
