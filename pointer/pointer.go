//pointers
/* pointers are used to store the memory address of another variable.*/
package main

import "fmt"

func main() {

	//a, b := 1, 2
	var a, b int
	fmt.Printf("Enter the 'a' and 'b' value: ")
	fmt.Scanf("%d%d", &a, &b)
	pptr(&a)

	// Task 1: Modify Total Calculation
	add, sub := ptr(&a, &b)
	fmt.Println("Addition Total:", add)
	fmt.Println("Sub Total:", sub)

}

func ptr(a, b *int) (int, int) {
	var add_total, sub_total int
	add_total = *a + *b
	sub_total = *a - *b

	return add_total, sub_total

}
func pptr(a *int) {
	if *a == 10 {
		*a = 20
	}

}
