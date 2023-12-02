package main

import (
	"fmt"
	"math"
	"strings"
)

func getInitials(name string) (string, string) {
	name = strings.ToUpper(name)
	var initials []string

	names := strings.Split(name, " ")

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

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
	learn_loops()
	names := []string{"Adam", "Eva", "Lucifer"}
	cycleNames(names, sayHello)

	area := circleArea(1)
	fmt.Println("Area of circle is", area)

	fn, ln := getInitials("lucifer morningstar")
	fmt.Println(fn, ln)
	fn, ln = getInitials("lucifer")
	fmt.Println(fn, ln)
}
