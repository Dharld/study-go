package models

import "fmt"

type Student struct {
	Name string 
	Age int32
	Grades []float32 
}

// Constructor remains similar
func NewStudent(name string, age int32) *Student {
    return &Student{
        Name: name,
        Age: age,
		Grades: []float32{},
    }
}

// Other methods
func (s *Student) AddGrade(grade float32) {
	s.Grades = append(s.Grades, grade)
}

func (s *Student) AverageGrade() float32 {
    if len(s.Grades) == 0 {
        return 0
    }
    
    var sum float32
    for _, grade := range s.Grades {
        sum += grade
    }
    return sum / float32(len(s.Grades))
}



func (s *Student) Display() {
    fmt.Printf("Name: %s, Age: %d, Grades: %v\n", s.Name, s.Age, s.Grades)
}