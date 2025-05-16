package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./students.db")
	if err != nil {
		return err
	}

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS students (name TEXT, age INTEGER, grade REAL)")
	return err
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func InsertStudent(student Student) error {
	_, err := DB.Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", student.Name, student.Age, student.Grade)
	return err
}

func GetStudents() ([]Student, error) {

	rows, err := DB.Query("SELECT name, age, grade FROM students")
	if err != nil {
		return nil, err
	}

	var students []Student
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Name, &student.Age, &student.Grade); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
