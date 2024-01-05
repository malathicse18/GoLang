/*
Mini Assignment: Slice Manipulation
Write a Go program that performs the following operations on slices:

Create a slice of integers with values from 1 to 5.
Append the values 6 and 7 to the slice.
Print the length and capacity of the slice.
Create a new slice by slicing the original slice from index 2 to 4 (inclusive).
Print the elements of the new slice.
*/
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5} //slice creation
	s = append(s, 6, 7)       //append slice
	fmt.Printf("length of the slice is %d and capacity of the slice is %d \n", len(s), cap(s))
	new_slice := s[2:4]					
	fmt.Printf("new slice is %v\n", new_slice)
	fmt.Printf("length of the slice is %d and capacity of the slice is %d \n", len(new_slice), cap(new_slice))

}
