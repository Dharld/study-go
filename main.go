package main

import (
	"fmt"

	models "study-go/internal/models"
	repository "study-go/internal/repository"
)

type StudentModel = models.Student

func main() {
	// Call the function
	userRepository := repository.NewUserRepository("users.txt")

	student := models.NewStudent("Yann 4", 16)
	userRepository.Create(*student)

	// Read from the repository
	students, err := userRepository.Read()
	if err != nil {
		fmt.Errorf("Error while reading the file")
	}

	fmt.Printf("Students %v\n", students)
}