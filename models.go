package main

import "github.com/arduclown/enternship-practise/utils"

type ResponseAge struct {
	Students []utils.Student `json: "students"`
	AverAge  float64         `json: "average_age"`
	MaxAge   int             `json: "max_age"`
}

type ResponseGrade struct {
	Students  []utils.Student `json: "students"`
	AverGrade float64         `json: "average_grade"`
	MaxGrade  float64         `json: "max_grade"`
}
