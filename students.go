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

func main() {
	baseStud := []Student{
		{Name: "Ann", Age: 16, AverageScore: 3.85},
		{Name: "Anton", Age: 16, AverageScore: 4},
		{Name: "Bob", Age: 17, AverageScore: 5},
		{Name: "Sarah", Age: 16, AverageScore: 4.9},
	}
	sort.Slice(baseStud, func(i, j int) bool {
		return baseStud[i].AverageScore > baseStud[j].AverageScore
	})

	for i, student := range baseStud {
		fmt.Printf("Student %d: Name=%s, Age=%d, AverageScore=%.2f\n",
			i+1, student.Name, student.Age, student.AverageScore)
	}

}
