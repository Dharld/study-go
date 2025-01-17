package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Utility to clear the screen
func clearScreen() {
    fmt.Print("\033[H\033[2J")
}

// Input and UI Utility Methods
func getInput(prompt string) string {
    fmt.Print(prompt)
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}

func pressEnterToContinue() {
    fmt.Println("\nPress Enter to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// Menu Method
func RunMenu(s StudentService) {
    for {
        clearScreen()
        fmt.Println("=== Student Management System ===")
        fmt.Println("1. Create Student")
        fmt.Println("2. List Students")
        fmt.Println("3. Exit")

        choice := getInput("Enter your choice: ")

        switch choice {
        case "1":
            s.CreateStudent()
        case "2":
            s.ListStudents()
        case "3":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
            pressEnterToContinue()
        }
    }
}