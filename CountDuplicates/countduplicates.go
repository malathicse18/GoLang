//Count Duplicates

package main

import "fmt"

func main() {
	a := []int{10, 20, 10, 40, 60, 20}
	m := make(map[int]int)

	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}
	fmt.Println(m)
}
