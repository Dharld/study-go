package repository

import (
	"io"
	"os"
	utils "study-go/internal/utils"
	"sync"

	models "study-go/internal/models"
)

type StudentModel = models.Student

type StudentRepository struct {
	filename string
	mu sync.Mutex
}

func NewUserRepository(filename string) *StudentRepository {
	return &StudentRepository{filename: filename}
}

func (r *StudentRepository) Read() ([]StudentModel, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

	filename := r.filename

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []StudentModel{}, nil
	}
   
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var students []StudentModel

	// Read until EOF
	for {
		// Read name
		name, err := utils.ReadString(file)
		if err != nil {
			// Check if we've reached the end of the file
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// Read age
		age, err := utils.ReadInt32(file)
		if err != nil {
			return nil, err
		}

		// Read grades
		grades, err := utils.ReadSlice[float32](file)

		student := StudentModel{
			Name: name,
			Age: age,
			Grades: grades,
		}

		students = append(students, student)
	}

	return students, nil
}

func (r *StudentRepository) Create(student StudentModel) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Open file in append mode
	file, err := os.OpenFile(r.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the name
	err = utils.WriteString(file, student.Name)
	if err != nil {
		return err
	}

	// Write age
	err = utils.WriteInt32(file, student.Age)
	if err != nil {
		return err
	}

	// Write grades 
    err = utils.WriteSlice(file, student.Grades)
	if err != nil {
		return err
	}

    return nil
}