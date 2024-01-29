//Find duplicates
package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 20, 30, 40, 50}
	l := len(a)
	m := make(map[int]bool)
	for i := 0; i < l; i++ {
		if m[a[i]] {
			fmt.Println("Duplicates:", a[i])
		}
		m[a[i]] = true
	}
}
