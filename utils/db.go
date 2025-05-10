package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS students (name TEXT, age INTEGER, grade REAL)")
	return err
}

func InsertStudent(student Student) error {
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", student.Name, student.Age, student.Grade)
	return err
}

func GetStudents() ([]Student, error) {
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, age, grade FROM students")
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
