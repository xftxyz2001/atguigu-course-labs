package model

import "fmt"

type student struct {
	Name string
	Age  int
}

func (s *student) String() string {
	return fmt.Sprintf("学生的名字是=%v, 年龄是=%v", s.Name, s.Age)
}

func NewStudent(name string, age int) *student {
	return &student{
		Name: name,
		Age:  age,
	}
}
