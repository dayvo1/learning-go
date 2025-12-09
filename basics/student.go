package main

import (
	"fmt"
)

type Student struct {
	Name string
	Id   string
	Gpa  float32
}

func (s *Student) UpdateGpa(newGpa float32) {
	s.Gpa = newGpa
}

func (s Student) Display() {
	fmt.Printf("Student Name: %s, Student ID: %s, Student GPA: %v\n", s.Name, s.Id, s.Gpa)
}
