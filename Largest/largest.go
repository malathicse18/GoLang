// Find Third Largest
package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{100, 1, 10, 20, 10, 40, 60, 20}
	first_max := math.MinInt
	second_max := math.MinInt
	third_max := math.MinInt

	for _, val := range a {
		if val > first_max {
			third_max = second_max
			second_max = first_max
			first_max = val
		} else if second_max < val && first_max > val {
			third_max = second_max
			second_max = val
		} else if third_max < val && second_max > val {

			third_max = val
		}
	}
	fmt.Println("First Max:", first_max)
	fmt.Println("Second Max:", second_max)
	fmt.Println("Third Max:", third_max)

}
