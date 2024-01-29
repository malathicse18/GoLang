// Array Reverse program
package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[i-l-1] = a[i-l-1], a[i]
	}
	fmt.Println("Array reversed:", a)
}
