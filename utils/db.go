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
