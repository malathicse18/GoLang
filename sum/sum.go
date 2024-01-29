package main

import "fmt"

func main() {
	a := []int{10, 20, 30}
	sum := 0
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	fmt.Print(sum)
}
