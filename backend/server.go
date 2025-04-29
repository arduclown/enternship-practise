// маршруты
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string  `json: "name"`
	Age   int     `json: "age"`
	Grade float64 `json: "grade"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func student(w http.ResponseWriter, req *http.Request) {
	s := Student{Name: "Alice", Age: 20, Grade: 4.5}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/student", student)

	http.ListenAndServe(":8090", nil)
}
