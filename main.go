package main

import (
	models "study-go/internal/models"
	repository "study-go/internal/repository"
	services "study-go/internal/services"
)

type StudentModel = models.Student

func main() {
	repo := repository.NewUserRepository("students.bin")
	studentService := services.NewStudentService(repo)

	// Run the menu
	services.RunMenu(*studentService)
}