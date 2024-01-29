package main

import "fmt"

func main() {
	var name = "Bharani"
	var age = 20
	fmt.Printf("Before update name is %v and age is %v", name, age)
	update(&name, &age)
	fmt.Printf("After update name is %v and age is %v", name, age)

}

func update(name *string, age *int) {
	*name = "Poovarasan"
	*age = 21
}
