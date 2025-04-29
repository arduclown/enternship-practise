// работа со структурами и указателями
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name         string
	Age          int
	AverageScore float32
}

func ChangeAge(s *Student, newAge int) {
	s.Age = newAge
}

func (s *Student) UpdateGrade(grade float32) {
	s.AverageScore = grade
}

func HighesGrade(s []Student) Student {
	if len(s) == 0 {
		return Student{}
	}

	maxS := s[0]
	for _, student := range s {
		if student.AverageScore > maxS.AverageScore {
			maxS = student
		}
	}

	return maxS
}

func SwapAges(s1, s2 *Student) {
	s1.Age, s2.Age = s2.Age, s1.Age
}

func main() {
	baseStud := []Student{
		{Name: "Ann", Age: 16, AverageScore: 3.85},
		{Name: "Anton", Age: 16, AverageScore: 4},
		{Name: "Bob", Age: 17, AverageScore: 5},
		{Name: "Sarah", Age: 16, AverageScore: 4.9},
	}
	sort.Slice(baseStud, func(i, j int) bool {
		return baseStud[i].AverageScore < baseStud[j].AverageScore
	})

	for i, student := range baseStud {
		fmt.Printf("Student %d: Name=%s, Age=%d, AverageScore=%.2f\n",
			i+1, student.Name, student.Age, student.AverageScore)
	}

	topStudent := HighesGrade(baseStud)
	fmt.Printf("Top student: %s. Average Score: %.2f\n", topStudent.Name, topStudent.AverageScore)

	SwapAges(&baseStud[0], &baseStud[3])
	for i, student := range baseStud {
		fmt.Printf("Student %d: Name=%s, Age=%d, AverageScore=%.2f\n",
			i+1, student.Name, student.Age, student.AverageScore)
	}
}
