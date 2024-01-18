// String Reverse program
package main

import "fmt"

func main() {
	//s := "Apple"
	var s string
	fmt.Print("Enter you string:")
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("Enter the string properly")
		return
	}
	reversed := stringReverse(s)
	fmt.Println("Reversed string:", reversed)
	palindrome := palindromeCheck(s, reversed)
	fmt.Println("The given string is:", palindrome)

}
func stringReverse(s string) string {
	var r string
	for i := 0; i < len(s); i++ {
		r = string(s[i]) + r

	}
	return r

}

//check the string is palindrome or not
func palindromeCheck(s, reversed string) string {
	if s == reversed {
		return "palindrome"
	} else {
		return "Not palindrome"
	}
}
