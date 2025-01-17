package models

type UndergraduateStudent struct {
	Student             `json:"student"`
	ProjectGrade string `json:"project_grade"`
}

// Other methods
func (g *UndergraduateStudent) GetAverageGrade() float32 {
	// Override the method
	return g.Student.AverageGrade() + 0.0125
}

