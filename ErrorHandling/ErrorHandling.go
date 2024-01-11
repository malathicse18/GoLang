package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter any number:")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(number)
	}

}
