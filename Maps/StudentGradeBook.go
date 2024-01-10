/*Write a Go program that manages a student grade book.
 The program should allow the user to perform the following actions:

1.Add a student and their corresponding grades to the grade book.
2.View the grades of a specific student.
3.Calculate and display the average grade for each student.
4.Display the student with the highest average grade.
5.Exit the program.*/

package main

import (
	"fmt"
)

func main() {
	gradeBook := make(map[string][]int)

	for {
		fmt.Println("\nGrade Book Menu:")
		fmt.Println("1. Add a student and their grades")
		fmt.Println("2. View grades of a specific student")
		fmt.Println("3. Calculate and display average grades for each student")
		fmt.Println("4. Display the student with the highest average grade")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice (1-5): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudent(gradeBook)
		case 2:
			viewStudentGrades(gradeBook)
		case 3:
			calculateAverageGrades(gradeBook)
		case 4:
			displayHighestAverageGradeStudent(gradeBook)
		case 5:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}

func addStudent(gradeBook map[string][]int) {
	var name string
	var grades []int

	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter grades (comma-separated): ")
	_, err := fmt.Scanf("%d,%d,%d,%d,%d", &grades[0], &grades[1], &grades[2], &grades[3], &grades[4])
	if err != nil {
		fmt.Println("Invalid input. Please enter numeric grades.")
		return
	}

	gradeBook[name] = grades
	fmt.Println("Student and grades added successfully.")
}

func viewStudentGrades(gradeBook map[string][]int) {
	var name string

	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	grades, exists := gradeBook[name]
	if exists {
		fmt.Printf("Grades of %s: %v\n", name, grades)
	} else {
		fmt.Println("Student not found.")
	}
}

func calculateAverageGrades(gradeBook map[string][]int) {
	for name, grades := range gradeBook {
		average := calculateAverage(grades)
		fmt.Printf("Average grade of %s: %.2f\n", name, average)
	}
}

func calculateAverage(grades []int) float64 {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}

func displayHighestAverageGradeStudent(gradeBook map[string][]int) {
	if len(gradeBook) == 0 {
		fmt.Println("No students in the grade book.")
		return
	}

	var highestAverageStudent string
	highestAverage := 0.0

	for name, grades := range gradeBook {
		average := calculateAverage(grades)
		if average > highestAverage {
			highestAverage = average
			highestAverageStudent = name
		}
	}

	fmt.Printf("Student with the highest average grade: %s\n", highestAverageStudent)
}
