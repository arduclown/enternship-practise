// работа со структурами и указателями
package main

type Student struct {
	Name  string  `json: "name"`
	Age   int     `json: "age"`
	Grade float64 `json: "grade"`
}

func ChangeAge(s *Student, newAge int) {
	s.Age = newAge
}

func (s *Student) UpdateGrade(gr float64) {
	s.Grade = gr
}

func HighesGrade(s []Student) Student {
	if len(s) == 0 {
		return Student{}
	}

	maxS := s[0]
	for _, student := range s {
		if student.Grade > maxS.Grade {
			maxS = student
		}
	}

	return maxS
}

func SwapAges(s1, s2 *Student) {
	s1.Age, s2.Age = s2.Age, s1.Age
}
