package models

type GraduateStudent struct {
	Student           
	ThesisGrade string
}


// Other methods
func (g *GraduateStudent) GetAverageGrade() float32 {
	// Override the method
	return g.Student.AverageGrade() + 0.125
}

