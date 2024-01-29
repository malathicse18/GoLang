// sort program using Bubble sort
package main

import "fmt"

func main() {
	a := []int{45, 20, 34, 90, 12}
	var l = len(a)
	isSwapped := true
	for isSwapped {
		isSwapped = false
		for i := 1; i < l; i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]

				isSwapped = true
			}

		}
		l--

	}
	fmt.Println(a)
}
