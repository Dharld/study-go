package service

import (
	"fmt"
	"strconv"
	"strings"

	models "study-go/internal/models"
	repository "study-go/internal/repository"
)

type StudentService struct {
    repo *repository.StudentRepository
}


func NewStudentService(repo *repository.StudentRepository) *StudentService {
    return &StudentService{repo: repo}
}






// Business Logic Methods
func (s *StudentService) CreateStudent() {
    clearScreen()
    fmt.Println("=== Create New Student ===")

    name := getInput("Enter student name: ")
    
    ageStr := getInput("Enter student age: ")
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        fmt.Println("Invalid age. Please enter a number.")
        pressEnterToContinue()
        return
    }

    gradesStr := getInput("Enter grades (comma-separated): ")
    gradeStrings := strings.Split(gradesStr, ",")
    grades := make([]float32, 0, len(gradeStrings))

    for _, gradeStr := range gradeStrings {
        grade, err := strconv.ParseFloat(strings.TrimSpace(gradeStr), 32)
        if err != nil {
            fmt.Println("Invalid grade. Skipping:", gradeStr)
            continue
        }
        grades = append(grades, float32(grade))
    }

    student := models.Student{
        Name:   name,
        Age:    int32(age),
        Grades: grades,
    }

    err = s.repo.Create(student)
    if err != nil {
        fmt.Println("Error creating student:", err)
    } else {
        fmt.Println("Student created successfully!")
    }

    pressEnterToContinue()
}

func (s *StudentService) ListStudents() {
    clearScreen()
    fmt.Println("=== Student List ===")

    students, err := s.repo.Read()

    if err != nil {
        fmt.Println("Error reading students:", err)
        pressEnterToContinue()
        return
    }

    if len(students) == 0 {
        fmt.Println("No students found.")
    } else {
        for i, student := range students {
            fmt.Printf("%d. Name: %s, Age: %d, Grades: %v\n", 
                i+1, student.Name, student.Age, student.Grades)
        }
    }

    pressEnterToContinue()
}

