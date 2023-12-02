package main

import (
	"fmt"
)

func learn_array() {
	// Array
	var ages []int = []int{1, 2, 3, 4, 5}
	fmt.Println("Created an array ages with a length of", len(ages))

	names := []string{"adam", "eva", "lucifer"}
	fmt.Println("New array of names", names)

	fmt.Println("Slice ages from index 1 to 3", ages[1:4])

	ages = append(ages, 9)
	fmt.Println("Appended ages array is", ages)
}

func main() {
	learn_array()

}
