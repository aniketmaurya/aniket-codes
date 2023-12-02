package main

import (
	"fmt"
	"math"
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

func learn_loops() {

	names := []string{"adam", "eva", "lucifer"}
	fmt.Println("New array of names", names)
	// Loop 5 times
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	// For in loop
	for index, value := range names {
		if index == 2 {
			fmt.Println("name at index 2 is", value)
		}
	}

}

func cycleNames(names []string, f func(string)) {
	for _, value := range names {
		f(value)
	}
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func circleArea(r float32) float32 {
	return math.Pi * r * r
}

func main() {
	sayHello("world")
	// learn_array()
	// learn_loops()
	// names := []string{"Adam", "Eva", "Lucifer"}
	// cycleNames(names, sayHello)

	// area := circleArea(1)
	// fmt.Println("Area of circle is", area)

	// fn, ln := getInitials("lucifer morningstar")
	// fmt.Println(fn, ln)
	// fn, ln = getInitials("lucifer")
	// fmt.Println(fn, ln)

	learn_map()

}
