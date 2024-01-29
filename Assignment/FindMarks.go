package main

import "fmt"

func main() {
	gradeBook := make(map[string]int)
	AddStudent(gradeBook)
	fmt.Println("Student Details:", gradeBook)
	marks := FindMarks(gradeBook)
	fmt.Println("75 scored students marks is:", marks)
}
func AddStudent(student map[string]int) {
	var name string
	var marks int
	for i := 1; i <= 5; i++ {
		fmt.Print("Enter student name: ")
		fmt.Scanln(&name)
		fmt.Print("Enter student mark: ")
		fmt.Scanln(&marks)
		student[name] = marks
	}

}
func FindMarks(grade map[string]int) []string {
	var neededMarks int
	fmt.Print("Enter the mark:")
	fmt.Scan(&neededMarks)
	var result []string
	for name, marks := range grade {
		if marks == neededMarks {
			result = append(result, fmt.Sprintf("%s:%d", name, marks))
		}
	}
	return result
}

//Another method
/*
package main

import "fmt"

type Student struct {
	name  string
	marks int
}

func main() {
	a := []Student{
		{"a", 75},
		{"b", 80},
		{"c", 90},
		{"d", 75},
	}

	result := FindMarks(a)
	fmt.Println(result)
}

func FindMarks(a []Student) map[int][]string {
	m := make(map[int][]string)

	for _, student := range a {
		if student.marks == 75 {
			m[student.marks] = append(m[student.marks], student.name)
		}
	}

	return m
}

*/
