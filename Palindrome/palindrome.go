package main

import (
	"fmt"
)

func main() {
	var name = "noon"
	isPalindrome := palindrome(name)

	if isPalindrome {
		fmt.Println("The given name is a palindrome.")
	} else {
		fmt.Println("The given name is not a palindrome.")
	}
}

func palindrome(name string) bool {
	for i := 0; i < len(name)/2; i++ {
		if name[i] != name[len(name)-i-1] {
			return false
		}
	}
	return true
}
