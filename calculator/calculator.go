package main

import "fmt"

func main() {
	var a, b int
	var choice int

	menu()
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scanln(&choice)
	fmt.Print("Enter first number:")
	fmt.Scanln(&a)
	fmt.Print("Enter Second number:")
	fmt.Scanln(&b)

	result := ArithmeticOperations(a, b, choice)
	fmt.Println("Result is:", result)

}

func menu() {
	var p = fmt.Println
	p("Calculator")
	p("1. Addition")
	p("2. Subtraction")
	p("3. Multiplication")
	p("4. Division")
}

func ArithmeticOperations(num1, num2 int, choice int) float64 {
	switch choice {
	case 1:
		return float64(num1 + num2)
	case 2:
		return float64(num1 - num2)
	case 3:
		return float64(num1 * num2)
	case 4:
		if num2 != 0 {
			return float64(num1) / float64(num2)
		} else {
			fmt.Println("Error: Cannot divide by zero.")
			return 0
		}
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		return 0
	}
}
