// Find which two number will become 60
package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 20, 30, 40, 50}
	m := make(map[int]bool)
	for _, val := range a {
		if m[60-val] {
			fmt.Println(val, 60-val)
		}
		m[val] = true

	}

}
